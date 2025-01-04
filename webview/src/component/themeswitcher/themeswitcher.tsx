import { Button } from "@nextui-org/react";
import { useTheme } from "next-themes";
import { useEffect, useState } from "react";
import { IoMdSunny, IoIosMoon } from "react-icons/io";

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
        <div>
            <Button
                variant='light'
                isIconOnly
                disableRipple
                disableAnimation
                className='button-setting'
                onClick={switchTheme}
            >
                {resolvedTheme === 'dark' ? (
                    <IoMdSunny className='h-full w-full' />
                ) : (
                    <IoIosMoon className='h-full w-full' />
                )}
            </Button>
        </div>
    );
}