import { Drawer, DrawerContent, DrawerHeader, DrawerBody, DrawerFooter, Input, Button } from "@nextui-org/react";
import { IoIosArrowRoundForward } from "react-icons/io";

interface ForwardProps {
    isOpen: boolean;
    setIsOpen: (isOpen: boolean) => void;
}

export function ForwardDrawer({ isOpen, setIsOpen }: ForwardProps) {
    return (<Drawer isOpen={isOpen} onOpenChange={setIsOpen}>
        <DrawerContent>
            {(onClose) => (
                <>
                    <DrawerHeader className="flex flex-col gap-1">Port Forwarding</DrawerHeader>
                    <DrawerBody>
                        <p>
                            It will forward the port of the device to the local port.
                        </p>
                        <div className="flex flex-row gap-2 items-center relative">
                            <Input label="Device Port" defaultValue="8080" />
                            <div className="h-full w-32 relative overflow-hidden">
                                <div className="arrows-container h-full w-fit">
                                    <IoIosArrowRoundForward className="h-full w-fit" />
                                    <IoIosArrowRoundForward className="h-full w-fit" />
                                    <IoIosArrowRoundForward className="h-full w-fit" />
                                </div>
                            </div>
                            <Input label="Local Port" defaultValue="8080" />
                        </div>
                    </DrawerBody>
                    <DrawerFooter>
                        <Button color="primary" onPress={Forward}>
                            Forward
                        </Button>
                    </DrawerFooter>
                </>
            )}
        </DrawerContent>
    </Drawer>)
}

function Forward() {
    console.log("Forward")
}