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
    const [width] = useState<number>(1000);
    const resizerRef = useRef<HTMLDivElement>(null);
    const terminalInstanceRef = useRef<{ term: XTerm, fitAddon: FitAddon } | null>(null);

    const writePrompt = (term: XTerm) => {
        term.write('\r\n$ ');
    };

    const executeCommand = useCallback(async (command: string, term: XTerm) => {
        if (!command) {
            writePrompt(term);
            return;
        }

        if (command === 'clear') {
            term.clear();
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
    }, []);

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

        // Wait for the DOM to be ready
        requestAnimationFrame(() => {
            if (terminalRef.current) {
                term.open(terminalRef.current);
                // Wait for terminal to be mounted
                requestAnimationFrame(() => {
                    try {
                        fitAddon.fit();
                        term.writeln('Terminal ready');
                        writePrompt(term);
                        terminalInstanceRef.current = { term, fitAddon };
                    } catch (e) {
                        console.error('Failed to initialize terminal:', e);
                    }
                });
            }
        });

        term.onData((data) => {
            const code = data.charCodeAt(0);
            if (code === 13) { // Enter
                const line = currentLineRef.current;
                executeCommand(line.trim(), term);
                currentLineRef.current = '';
                return;
            }

            if (code === 127 || code === 8) { // Backspace
                if (currentLineRef.current.length > 0) {
                    currentLineRef.current = currentLineRef.current.slice(0, -1);
                    term.write('\b \b');
                }
                return;
            }

            if (code >= 32) { // Printable characters
                currentLineRef.current += data;
                term.write(data);
            }
        });

        const handleResize = () => {
            if (terminalInstanceRef.current) {
                requestAnimationFrame(() => {
                    try {
                        terminalInstanceRef.current?.fitAddon.fit();
                    } catch (e) {
                        console.error('Failed to fit terminal:', e);
                    }
                });
            }
        };

        window.addEventListener('resize', handleResize);
        return () => {
            term.dispose();
            window.removeEventListener('resize', handleResize);
        };
    }, [executeCommand]);

    return (
        <div className="terminal-container" style={{ width: `${width}px` }}>
            <div ref={terminalRef} className="terminal" />
            <div ref={resizerRef} className="terminal-resizer" />
        </div>
    );
}
