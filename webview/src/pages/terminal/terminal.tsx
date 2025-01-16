import { useEffect, useRef, useState } from 'react';
import { Terminal as XTerm } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
import { WebLinksAddon } from 'xterm-addon-web-links';
import { SearchAddon } from 'xterm-addon-search';
import 'xterm/css/xterm.css';
import './terminal.css';
import { CommandResponse } from 'src/proto/proto/shifu/shifu_pb';
import { ExecuteCommand, GetCompletions } from 'src/apis/shifu/terminal';

export default function Terminal() {
    const terminalRef = useRef<HTMLDivElement>(null);
    const [terminal, setTerminal] = useState<XTerm | null>(null);
    const commandHistoryRef = useRef<string[]>([]);
    const historyIndexRef = useRef<number>(-1);
    const currentLineRef = useRef<string>('');
    const cursorPositionRef = useRef<number>(0);
    const [width, setWidth] = useState<number>(1000);
    const resizerRef = useRef<HTMLDivElement>(null);
    const isResizingRef = useRef(false);
    const terminalInstanceRef = useRef<{ term: XTerm, fitAddon: FitAddon } | null>(null);
    const [completions, setCompletions] = useState<string[]>([]);
    const [showCompletions, setShowCompletions] = useState(false);
    const [selectedCompletion, setSelectedCompletion] = useState(0);
    const [currentDir, setCurrentDir] = useState<string>(process.env.HOME || '/');
    const selectedCompletionRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        if (!terminalRef.current) return;

        const initTerminal = () => {
            const term = new XTerm({
                cursorBlink: true,
                fontSize: 14,
                fontFamily: 'Menlo, Monaco, "Courier New", monospace',
                theme: {
                    background: '#1e1e1e',
                    foreground: '#ffffff'
                },
                allowTransparency: true,
                convertEol: true,
                cursorStyle: 'block',
                scrollback: 1000,
                cols: 80,
                rows: 24
            });

            const fitAddon = new FitAddon();
            const webLinksAddon = new WebLinksAddon();
            const searchAddon = new SearchAddon();

            term.loadAddon(fitAddon);
            term.loadAddon(webLinksAddon);
            term.loadAddon(searchAddon);

            if (terminalRef.current?.offsetHeight && terminalRef.current?.offsetWidth) {
                term.open(terminalRef.current);

                setTimeout(() => {
                    try {
                        fitAddon.fit();
                        term.writeln('You can execute local commands here');
                        term.writeln('Have a good day! 😋');
                        writePrompt(term);
                        setTerminal(term);
                    } catch (e) {
                        console.error('Failed to initialize terminal:', e);
                    }
                }, 100);
            }

            term.onData((data) => {
                const code = data.charCodeAt(0);
                if (code === 22) { // Ctrl+V
                    navigator.clipboard.readText().then((text) => {
                        handlePaste(text, term);
                    });
                    return;
                }
                handleInput(data, term);
            });

            term.onKey(({ key, domEvent }) => {
                const ev = domEvent as KeyboardEvent;

                if (ev.ctrlKey) {
                    switch (ev.code) {
                        case 'KeyC':
                            if (terminal?.hasSelection()) {
                                document.execCommand('copy');
                            } else {
                                handleCtrlC(term);
                            }
                            return;
                        case 'KeyD':
                            handleCtrlD(term);
                            return;
                    }
                }

                if (ev.key === 'ArrowUp') {
                    handleArrowUp(term);
                    return;
                }
                if (ev.key === 'ArrowDown') {
                    handleArrowDown(term);
                    return;
                }
                if (ev.key === 'ArrowLeft') {
                    handleArrowLeft(term);
                    return;
                }
                if (ev.key === 'ArrowRight') {
                    handleArrowRight(term);
                    return;
                }
            });

            const handleResize = () => {
                setTimeout(() => {
                    try {
                        fitAddon.fit();
                    } catch (e) {
                        console.error('Failed to fit terminal:', e);
                    }
                }, 0);
            };

            handleResize();
            window.addEventListener('resize', handleResize);

            terminalInstanceRef.current = { term, fitAddon };

            return () => {
                term.dispose();
                window.removeEventListener('resize', handleResize);
            };
        };

        const timeoutId = setTimeout(initTerminal, 0);
        return () => clearTimeout(timeoutId);
    }, []);

    useEffect(() => {
        const handleMouseDown = (e: MouseEvent) => {
            isResizingRef.current = true;
            document.addEventListener('mousemove', handleMouseMove);
            document.addEventListener('mouseup', handleMouseUp);
        };

        const handleMouseMove = (e: MouseEvent) => {
            if (!isResizingRef.current) return;

            const containerLeft = terminalRef.current?.getBoundingClientRect().left || 0;
            const newWidth = Math.max(200, Math.min(e.clientX - containerLeft, window.innerWidth - containerLeft));

            setWidth(newWidth);

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

        const handleMouseUp = () => {
            isResizingRef.current = false;
            document.removeEventListener('mousemove', handleMouseMove);
            document.removeEventListener('mouseup', handleMouseUp);
        };

        const resizer = resizerRef.current;
        if (resizer) {
            resizer.addEventListener('mousedown', handleMouseDown);
        }

        return () => {
            if (resizer) {
                resizer.removeEventListener('mousedown', handleMouseDown);
            }
            document.removeEventListener('mousemove', handleMouseMove);
            document.removeEventListener('mouseup', handleMouseUp);
        };
    }, []);

    const writePrompt = (term: XTerm) => {
        const shortDir = currentDir.replace(/^.*[/]/, '');
        term.write(`\r\n\x1b[32m${shortDir}\x1b[0m $ `);
    };

    const handleInput = (data: string, term: XTerm) => {
        const code = data.charCodeAt(0);
        const line = currentLineRef.current;
        const pos = cursorPositionRef.current;

        if (code === 9) { // Tab
            handleTab(term);
            return;
        }

        if (code === 13) { // Enter
            handleEnter(term, line);
            return;
        }

        if (showCompletions) {
            setShowCompletions(false);
            setCompletions([]);
        }

        if (code === 127 || code === 8) { // Backspace
            if (pos > 0) {
                const newLine = line.slice(0, pos - 1) + line.slice(pos);
                cursorPositionRef.current--;
                rewriteLine(term, newLine);
            }
        } else if (code >= 32) { // Printable characters
            const newLine = line.slice(0, pos) + data + line.slice(pos);
            cursorPositionRef.current++;
            rewriteLine(term, newLine);
        }
    };

    const handlePaste = (text: string, term: XTerm) => {
        const lines = text.split(/\r?\n/);
        lines.forEach((line, i) => {
            if (i === 0) {
                currentLineRef.current += line;
                term.write(line);
            } else {
                executeCommand(currentLineRef.current, term);
                currentLineRef.current = line;
                term.write(line);
            }
        });
    };

    const handleCtrlC = (term: XTerm) => {
        term.write('^C');
        currentLineRef.current = '';
        cursorPositionRef.current = 0;
        writePrompt(term);
    };

    const handleCtrlD = (term: XTerm) => {
        if (currentLineRef.current === '') {
            term.write('exit\r\n');
            // 可以在这里处理终端退出逻辑
        }
    };

    const handleArrowUp = (term: XTerm) => {
        if (showCompletions && completions.length > 0) {
            setSelectedCompletion(prev => {
                const newIndex = prev > 0 ? prev - 1 : completions.length - 1;
                const completion = completions[newIndex];
                const words = currentLineRef.current.split(' ');
                const beforeLastWord = words.slice(0, -1).join(' ');
                const newLine = beforeLastWord +
                    (beforeLastWord ? ' ' : '') +
                    completion;
                rewriteLine(term, newLine);
                return newIndex;
            });
            return;
        }

        const history = commandHistoryRef.current;
        if (history.length === 0) return;

        if (historyIndexRef.current < history.length - 1) {
            historyIndexRef.current++;
            const historyCommand = history[history.length - 1 - historyIndexRef.current];
            rewriteLine(term, historyCommand);
        }
    };

    const handleArrowDown = (term: XTerm) => {
        if (showCompletions && completions.length > 0) {
            setSelectedCompletion(prev => {
                const newIndex = prev < completions.length - 1 ? prev + 1 : 0;
                const completion = completions[newIndex];
                const words = currentLineRef.current.split(' ');
                const beforeLastWord = words.slice(0, -1).join(' ');
                const newLine = beforeLastWord +
                    (beforeLastWord ? ' ' : '') +
                    completion;
                rewriteLine(term, newLine);
                return newIndex;
            });
            return;
        }

        const history = commandHistoryRef.current;
        if (historyIndexRef.current > -1) {
            historyIndexRef.current--;
            const historyCommand = historyIndexRef.current === -1 ? ''
                : history[history.length - 1 - historyIndexRef.current];
            rewriteLine(term, historyCommand);
        }
    };

    const handleArrowLeft = (term: XTerm) => {
        if (cursorPositionRef.current > 0) {
            cursorPositionRef.current--;
            term.write('\b');
        }
    };

    const handleArrowRight = (term: XTerm) => {
        if (cursorPositionRef.current < currentLineRef.current.length) {
            term.write(currentLineRef.current[cursorPositionRef.current]);
            cursorPositionRef.current++;
        }
    };

    const rewriteLine = (term: XTerm, newLine?: string) => {
        const line = newLine ?? currentLineRef.current;
        const shortDir = currentDir.replace(/^.*[/]/, '');
        const prompt = `\x1b[32m${shortDir}\x1b[0m $ `;

        // 清除当前行并重写
        term.write('\x1b[2K\r');
        term.write(prompt + line);

        // 如果光标不在行尾，需要移动回正确位置
        if (cursorPositionRef.current < line.length) {
            term.write('\x1b[' + (prompt.length + cursorPositionRef.current) + 'G');
        }

        currentLineRef.current = line;
    };

    const handleEnter = (term: XTerm, line: string) => {
        if (showCompletions) {
            setShowCompletions(false);
            setCompletions([]);
            return;
        }

        executeCommand(line.trim(), term);
        if (line.trim()) {
            commandHistoryRef.current.push(line.trim());
        }
        historyIndexRef.current = -1;
        currentLineRef.current = '';
        cursorPositionRef.current = 0;
    };

    const executeCommand = async (command: string, term: XTerm) => {
        if (!command) {
            writePrompt(term);
            return;
        }

        if (command.startsWith('cd ')) {
            const newDir = command.slice(3).trim();
            try {
                const { stream } = ExecuteCommand(`cd "${newDir}" && pwd`);
                stream.on('data', (response: CommandResponse) => {
                    if (!response.getIsError()) {
                        const output = response.getOutput().trim();
                        if (output) {
                            setCurrentDir(output);
                        }
                    }
                });
                writePrompt(term);
                return;
            } catch (error) {
                term.write(`\r\n\x1b[31mError: ${error}\x1b[0m`);
                writePrompt(term);
                return;
            }
        }

        if (command.trim() === 'clear') {
            term.write('\x1b[2J');    // 清除整个屏幕
            term.write('\x1b[3J');    // 清除滚动缓冲区
            term.write('\x1b[H');     // 移动光标到顶部
            writePrompt(term);
            return;
        }

        try {
            const { stream } = ExecuteCommand(command);

            stream.on('data', (response: CommandResponse) => {
                const output = response.getOutput();
                const isError = response.getIsError();

                if (isError) {
                    term.write(`\r\n\x1b[31m${output}\x1b[0m`);
                } else {
                    term.write(`\r\n${output}`);
                }
            });

            stream.on('end', () => {
                writePrompt(term);
            });

            stream.on('error', (err) => {
                term.write(`\r\n\x1b[31mError: ${err.message}\x1b[0m`);
                writePrompt(term);
            });
        } catch (error) {
            term.write(`\r\n\x1b[31mError: ${error}\x1b[0m`);
            writePrompt(term);
        }
    };

    const handleTab = async (term: XTerm) => {
        const line = currentLineRef.current;
        const words = line.split(' ');
        const lastWord = words[words.length - 1];
        const beforeLastWord = words.slice(0, -1).join(' ');

        try {
            const suggestions = await GetCompletions(lastWord, currentDir);
            if (suggestions.length === 0) {
                return;
            }

            if (suggestions.length === 1) {
                // 直接补全
                const completion = suggestions[0];
                currentLineRef.current = beforeLastWord +
                    (beforeLastWord ? ' ' : '') +
                    completion;
                rewriteLine(term);
            } else {
                // 显示补全选项
                term.write('\r\n');
                const maxLength = Math.max(...suggestions.map(s => s.length));
                const termWidth = term.cols;
                const columns = Math.floor(termWidth / (maxLength + 2));
                const rows = Math.ceil(suggestions.length / columns);

                for (let row = 0; row < rows; row++) {
                    const rowItems = suggestions.slice(row * columns, (row + 1) * columns);
                    const rowText = rowItems.map(s => s.padEnd(maxLength + 2)).join('');
                    term.write(rowText + '\r\n');
                }

                setCompletions(suggestions);
                setShowCompletions(true);
                setSelectedCompletion(0);
                writePrompt(term);
                term.write(line);
            }
        } catch (error) {
            console.error('Completion error:', error);
        }
    };

    const handleSpecialKeys = (event: KeyboardEvent, term: XTerm) => {
        // Ctrl + A: 移动到行首
        if (event.ctrlKey && event.key === 'a') {
            cursorPositionRef.current = 0;
            rewriteLine(term);
            return true;
        }
        // Ctrl + E: 移动到行尾
        if (event.ctrlKey && event.key === 'e') {
            cursorPositionRef.current = currentLineRef.current.length;
            rewriteLine(term);
            return true;
        }
        // Ctrl + L: 清屏
        if (event.ctrlKey && event.key === 'l') {
            term.write('\x1b[2J\x1b[H');
            rewriteLine(term);
            return true;
        }
        // Home: 移动到行首
        if (event.key === 'Home') {
            cursorPositionRef.current = 0;
            rewriteLine(term);
            return true;
        }
        // End: 移动到行尾
        if (event.key === 'End') {
            cursorPositionRef.current = currentLineRef.current.length;
            rewriteLine(term);
            return true;
        }
        return false;
    };

    return (
        <div className="terminal-container" style={{ width: `${width}px` }}>
            <div ref={terminalRef} className="terminal" />
            <div ref={resizerRef} className="terminal-resizer" />
        </div>
    );
}
