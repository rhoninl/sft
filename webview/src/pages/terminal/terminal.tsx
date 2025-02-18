import { useEffect, useRef, useState, useCallback } from 'react';
import { Terminal as XTerm } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
import 'xterm/css/xterm.css';
import './terminal.css';
import { CommandResponse } from 'src/proto/proto/shifu/shifu_pb';
import { ExecuteCommand } from 'src/apis/shifu/terminal';

export default function Terminal() {
    const terminalRef = useRef<HTMLDivElement>(null);
    const currentLineRef = useRef<string>('');
    const commandHistoryRef = useRef<string[]>([]);
    const historyIndexRef = useRef<number>(0);
    const tempInputRef = useRef<string>('');
    const [width] = useState<number>(1000);
    const resizerRef = useRef<HTMLDivElement>(null);
    const terminalInstanceRef = useRef<{ term: XTerm, fitAddon: FitAddon } | null>(null);
    const searchModeRef = useRef<boolean>(false);
    const searchQueryRef = useRef<string>('');

    // Write prompt to the terminal
    const writePrompt = useCallback((term: XTerm) => {
        term.write('\r\n> ');
    }, []);

    const updateSearchDisplay = useCallback((term: XTerm) => {
        const query = searchQueryRef.current;
        let match = '';
        // Search in history from newest to oldest
        for (let i = commandHistoryRef.current.length - 1; i >= 0; i--) {
            const cmd = commandHistoryRef.current[i];
            if (cmd.includes(query)) {
                match = cmd;
                break;
            }
        }
        term.write('\r\x1b[Kreverse-i-search: ' + query + (match ? ' (match: ' + match + ')' : ''));
    }, []);

    // Execute command and handle stream events
    const executeCommand = useCallback(async (command: string, term: XTerm) => {
        if (!command) {
            writePrompt(term);
            return;
        }

        if (command === 'clear') {
            // Clear the screen and scrollback buffer
            term.write('\x1b[2J\x1b[3J\x1b[H');
            term.writeln('Terminal ready');
            writePrompt(term);
            return;
        }

        try {
            const { stream } = ExecuteCommand(command);
            stream.on('data', (response: CommandResponse) => {
                const output = response.getOutput();
                term.write('\r\n' + output);
            });
            stream.on('end', () => {
                writePrompt(term);
            });
            stream.on('error', (err) => {
                term.write(`\r\nError: ${err.message}`);
                writePrompt(term);
            });
        } catch (error) {
            term.write(`\r\nError: ${error}`);
            writePrompt(term);
        }
    }, [writePrompt]);

    // Handle terminal input data with command history support
    const handleInput = useCallback((data: string, term: XTerm) => {
        // Check for Ctrl+R to initiate reverse search (ctrl+R has char code 18)
        if (data.charCodeAt(0) === 18 && !searchModeRef.current) {
            searchModeRef.current = true;
            searchQueryRef.current = '';
            term.write('\r\n(reverse-i-search): ');
            return;
        }

        // If in reverse search mode, handle input differently
        if (searchModeRef.current) {
            const charCode = data.charCodeAt(0);
            if (charCode === 13) { // Enter: accept current match
                searchModeRef.current = false;
                let match = '';
                for (let i = commandHistoryRef.current.length - 1; i >= 0; i--) {
                    const cmd = commandHistoryRef.current[i];
                    if (cmd.includes(searchQueryRef.current)) {
                        match = cmd;
                        break;
                    }
                }
                term.write('\r\x1b[K');
                term.write('> ' + match);
                currentLineRef.current = match;
                return;
            } else if (charCode === 8 || charCode === 127) { // Backspace in search mode
                if (searchQueryRef.current.length > 0) {
                    searchQueryRef.current = searchQueryRef.current.slice(0, -1);
                }
                updateSearchDisplay(term);
                return;
            } else if (charCode === 3) { // Ctrl+C in search mode cancels search
                searchModeRef.current = false;
                searchQueryRef.current = '';
                term.write('\r\n^C');
                writePrompt(term);
                return;
            } else if (data === '\x1b[A' || data === '\x1b[B') {
                // Optionally handle arrow keys in search mode (not implemented here)
                return;
            } else if (charCode >= 32) { // Printable character in search mode
                searchQueryRef.current += data;
                updateSearchDisplay(term);
                return;
            }
        }

        // Regular input processing with history support
        if (data === '\x1b[A') { // Arrow Up
            if (commandHistoryRef.current.length > 0) {
                if (historyIndexRef.current === commandHistoryRef.current.length) {
                    tempInputRef.current = currentLineRef.current;
                }
                historyIndexRef.current = Math.max(0, historyIndexRef.current - 1);
                const histCommand = commandHistoryRef.current[historyIndexRef.current] || '';
                term.write('\r\x1b[K');
                term.write('> ' + histCommand);
                currentLineRef.current = histCommand;
            }
            return;
        } else if (data === '\x1b[B') { // Arrow Down
            if (commandHistoryRef.current.length > 0) {
                historyIndexRef.current = Math.min(commandHistoryRef.current.length, historyIndexRef.current + 1);
                let histCommand = '';
                if (historyIndexRef.current === commandHistoryRef.current.length) {
                    histCommand = tempInputRef.current;
                } else {
                    histCommand = commandHistoryRef.current[historyIndexRef.current] || '';
                }
                term.write('\r\x1b[K');
                term.write('> ' + histCommand);
                currentLineRef.current = histCommand;
            }
            return;
        }

        const charCode = data.charCodeAt(0);
        if (charCode === 3) { // Ctrl+C
            term.write('\r\n^C');
            currentLineRef.current = '';
            writePrompt(term);
            return;
        } else if (charCode === 13) { // Enter
            const command = currentLineRef.current.trim();
            if (command !== '') {
                commandHistoryRef.current.push(command);
            }
            historyIndexRef.current = commandHistoryRef.current.length;
            executeCommand(command, term);
            currentLineRef.current = '';
        } else if (charCode === 8 || charCode === 127) { // Backspace
            if (currentLineRef.current.length > 0) {
                currentLineRef.current = currentLineRef.current.slice(0, -1);
                term.write('\b \b');
            }
        } else if (charCode >= 32) { // Printable characters
            currentLineRef.current += data;
            term.write(data);
        }
    }, [executeCommand, writePrompt, updateSearchDisplay]);

    // Handle window resize to adjust terminal size
    const handleResize = useCallback(() => {
        if (terminalInstanceRef.current) {
            requestAnimationFrame(() => {
                try {
                    terminalInstanceRef.current?.fitAddon.fit();
                } catch (e) {
                    console.error('Failed to fit terminal:', e);
                }
            });
        }
    }, []);

    // Initialize the terminal when component mounts
    useEffect(() => {
        if (!terminalRef.current) return;

        const term = new XTerm({
            cursorBlink: true,
            fontSize: 14,
            fontFamily: 'Menlo, Monaco, "Courier New", monospace',
            theme: {
                background: '#1e1e1e',
                foreground: '#ffffff'
            },
            cols: 80,
            rows: 24
        });

        const fitAddon = new FitAddon();
        term.loadAddon(fitAddon);

        // Open the terminal
        term.open(terminalRef.current);

        // Fit terminal after mounting with a slight delay
        setTimeout(() => {
            try {
                fitAddon.fit();
                term.writeln('Terminal ready');
                writePrompt(term);
                terminalInstanceRef.current = { term, fitAddon };
            } catch (error) {
                console.error('Failed to initialize terminal:', error);
            }
        }, 100);

        // Set up input handler
        term.onData((data) => handleInput(data, term));

        // Listen for window resize events
        window.addEventListener('resize', handleResize);

        return () => {
            term.dispose();
            window.removeEventListener('resize', handleResize);
        };
    }, [handleInput, handleResize, writePrompt]);


    return (
        <div className="terminal-container" style={{ width: `${width}px`, height: '500px' }}>
            <div ref={terminalRef} className="terminal" style={{ height: '100%' }} />
            <div ref={resizerRef} className="terminal-resizer" />
        </div>
    );
}
