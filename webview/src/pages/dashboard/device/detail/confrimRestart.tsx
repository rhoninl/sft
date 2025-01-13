import { Button, ModalBody, Modal, ModalContent, ModalFooter, ModalHeader } from "@nextui-org/react";
import { RestartDevice } from "src/apis/shifu/device";

interface ConfirmDeleteProps {
    deviceName: string;
    isOpen: boolean;
    setIsOpen: (isOpen: boolean) => void;
}

export function ConfirmRestart({ deviceName, isOpen, setIsOpen }: ConfirmDeleteProps) {
    function restartDevice() {
        RestartDevice(deviceName).then(() => {
            setIsOpen(false);
        }).catch((error) => {
            console.error(error);
        });
    }
    return <>
        <Modal isOpen={isOpen} onOpenChange={setIsOpen}>
            <ModalContent>
                <ModalHeader>
                    <span className="text-yellow-500">Are you sure you want to restart?</span>
                </ModalHeader>
                <ModalBody>
                    <p> <span className="font-extrabold">{deviceName}</span> will be restarted!</p>
                </ModalBody>
                <ModalFooter>
                    <Button color="warning" onClick={() => restartDevice()} > Restart </Button>
                    <Button color="primary" onClick={() => setIsOpen(false)}> Cancel </Button>
                </ModalFooter>
            </ModalContent>
        </Modal >
    </>
}