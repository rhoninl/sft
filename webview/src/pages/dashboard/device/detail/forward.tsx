import { Drawer, DrawerContent, DrawerHeader, DrawerBody, DrawerFooter, Input, Button } from "@nextui-org/react";
import { IoIosArrowRoundForward } from "react-icons/io";
import { ForwardPort } from "src/apis/shifu/device";
import { useEffect, useState } from "react";
interface ForwardProps {
    deviceName: string;
    isOpen: boolean;
    setIsOpen: (isOpen: boolean) => void;
}

export function ForwardDrawer({ deviceName, isOpen, setIsOpen }: ForwardProps) {
    const [devicePort, setDevicePort] = useState("8080");
    const [localPort, setLocalPort] = useState("8080");
    const [isForwarding, setIsForwarding] = useState(false);
    const [cancel, setCancel] = useState<() => void>(() => { });

    function forward() {
        const { promise, cancel } = ForwardPort(deviceName, devicePort, localPort);
        promise.then((success) => {
            if (!success) {
                console.error("Failed to forward port");
                return
            }
            setCancel(() => cancel);
            setIsForwarding(success);
        }).catch((err) => {
            console.error("Failed to forward port:", err);
            setIsForwarding(false);
        });
    }

    const handleDrawerClose = () => {
        if (isForwarding) {
            cancel();
            setIsForwarding(false);
        }
        setIsOpen(false);
    };

    useEffect(() => {
        return () => {
            if (isForwarding) {
                console.log("Cancelling port forwarding");
                cancel();
            }
        }
    }, [isForwarding, cancel]);

    return (<Drawer
        isOpen={isOpen}
        onOpenChange={handleDrawerClose}
    >
        <DrawerContent>
            {(onClose) => (
                <>
                    <DrawerHeader className="flex flex-col gap-1">Port Forwarding</DrawerHeader>
                    <DrawerBody>
                        <p>
                            It will forward the port of the device to the local port.
                        </p>
                        <div className="flex flex-row gap-2 items-center relative">
                            <Input
                                type="number"
                                label="Device Port"
                                defaultValue="8080"
                                onChange={(e) => setDevicePort(e.target.value)}
                                onKeyDown={inputNumber}
                                min={1}
                                max={65535}
                            />
                            <div className="h-full w-32 relative overflow-hidden">
                                <div className={`arrows-container h-full w-fit ${isForwarding ? 'animate-scroll' : ''}`}>
                                    <IoIosArrowRoundForward className="h-full w-fit" />
                                    <IoIosArrowRoundForward className="h-full w-fit" />
                                    <IoIosArrowRoundForward className="h-full w-fit" />
                                </div>
                            </div>
                            <Input
                                type="number"
                                label="Local Port"
                                defaultValue="8080"
                                onChange={(e) => setLocalPort(e.target.value)}
                                onKeyDown={inputNumber}
                                min={1}
                                max={65535}
                            />
                        </div>
                    </DrawerBody>
                    <DrawerFooter>
                        <Button
                            color={isForwarding ? "danger" : "primary"}
                            onPress={() => {
                                if (isForwarding) {
                                    cancel();
                                    setIsForwarding(false);
                                } else {
                                    forward();
                                }
                            }}
                        >
                            {isForwarding ? "Stop" : "Forward"}
                        </Button>
                    </DrawerFooter>
                </>
            )}
        </DrawerContent>
    </Drawer>)
}

function inputNumber(e: React.KeyboardEvent<HTMLInputElement>) {
    const allowedKeys = [
        'Backspace',
        'Delete',
        'ArrowLeft',
        'ArrowRight',
        'ArrowUp',
        'ArrowDown',
        'Tab'
    ];

    if (allowedKeys.includes(e.key)) {
        return;
    }

    if (e.key === 'e' || e.key === '-') {
        e.preventDefault();
        return;
    }

    if (!e.key.match(/^[0-9]$/)) {
        e.preventDefault();
    }
}

