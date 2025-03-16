import React, { useEffect, useState, useRef } from 'react';
import confetti from 'canvas-confetti';
import { useTheme } from 'next-themes';

const konamiCode = ['s', 'h', 'i', 'f', 'u'];

const EasterEgg: React.FC = () => {
    const [keySequence, setKeySequence] = useState<string[]>([]);
    const [show, setShow] = useState<boolean>(false);
    const { theme, setTheme, themes } = useTheme();
    const [previousTheme, setPreviousTheme] = useState<string>('light');
    const themeTimeoutRef = useRef<ReturnType<typeof setTimeout> | null>(null);

    useEffect(() => {
        const handleKeyDown = (e: KeyboardEvent) => {
            setKeySequence(prev => {
                const newSeq = [...prev, e.key];
                if (newSeq.length > konamiCode.length) {
                    newSeq.shift();
                }
                if (newSeq.join('').toLowerCase() === konamiCode.join('').toLowerCase()) {
                    console.log('Easter egg activated!');
                    if (!show) {
                        setPreviousTheme((theme === 'system' || !theme) ? 'light' : theme);
                        setShow(true);
                    } else {
                        if (themeTimeoutRef.current) {
                            clearTimeout(themeTimeoutRef.current);
                            themeTimeoutRef.current = null;
                        }
                        document.documentElement.classList.remove('cyber');
                        setTheme('system');
                        console.log(themes)
                        setShow(false);
                    }
                }
                return newSeq;
            });
        };

        window.addEventListener('keydown', handleKeyDown);
        return () => window.removeEventListener('keydown', handleKeyDown);
    }, [show, theme, setTheme]);

    useEffect(() => {
        if (show) {
            const duration = 5000; // 5 seconds of fireworks
            const animationEnd = Date.now() + duration;

            const interval = setInterval(() => {
                const timeLeft = animationEnd - Date.now();
                if (timeLeft <= 0) {
                    clearInterval(interval);
                    return;
                }
                // Fire two bursts with random origins for a varied effect
                confetti({
                    particleCount: 50,
                    angle: 60,
                    spread: 55,
                    origin: { x: Math.random(), y: Math.random() - 0.2 }
                });
                confetti({
                    particleCount: 50,
                    angle: 120,
                    spread: 55,
                    origin: { x: Math.random(), y: Math.random() - 0.2 }
                });
            }, 250);

            return () => clearInterval(interval);
        }
    }, [show]);

    useEffect(() => {
        if (show) {
            const timeout = setTimeout(() => setShow(false), 5000);
            return () => clearTimeout(timeout);
        }
    }, [show]);

    useEffect(() => {
        if (show) {
            setTheme('cyber');
            themeTimeoutRef.current = setTimeout(() => {
                console.log('Restoring theme to:', previousTheme);
                setTheme(previousTheme);
                themeTimeoutRef.current = null;
            }, 5000);

            return () => {
                if (themeTimeoutRef.current) {
                    clearTimeout(themeTimeoutRef.current);
                    themeTimeoutRef.current = null;
                }
            };
        }
    }, [show, previousTheme, setTheme]);

    return null;
};

export default EasterEgg; 