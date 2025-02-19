import { Button } from "@heroui/react";
import { useTheme } from "next-themes";
import { useEffect, useState } from "react";
import { IoSunny, IoMoon } from "react-icons/io5";

export function ThemeSwitcher() {
    const [mounted, setMounted] = useState(false);
    const { resolvedTheme, setTheme } = useTheme();

    useEffect(() => {
        setMounted(true);
        if (!mounted) {
            setTheme("system")
        }
    }, [setTheme, mounted]);

    function switchTheme() {
        if (resolvedTheme === 'dark') {
            setTheme('light');
        } else {
            setTheme('dark');
        }
    }

    if (!mounted) return null;

    return (
        <Button
            variant='light'
            isIconOnly
            disableRipple
            disableAnimation
            className='button-setting'
            onClick={switchTheme}
        >
            {resolvedTheme === 'dark' ? (
                <IoSunny className='h-full w-full' />
            ) : (
                <IoMoon className='h-full w-full' />
            )}
        </Button>
    );
}