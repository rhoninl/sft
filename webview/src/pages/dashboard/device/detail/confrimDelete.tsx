import { Button, ModalBody, Modal, ModalContent, ModalFooter, ModalHeader, Input } from "@nextui-org/react";
import { useState, useEffect } from "react";
import { DeleteDevice } from "src/apis/shifu/device";

interface ConfirmDeleteProps {
    deviceName: string;
    isOpen: boolean;
    setIsOpen: (isOpen: boolean) => void;
}

export function ConfirmDelete({ deviceName, isOpen, setIsOpen }: ConfirmDeleteProps) {
    const [confirmation, setConfirmation] = useState("");

    useEffect(() => {
        setConfirmation("");
    }, [isOpen]);

    function deleteDevice() {
        if (confirmation.toLowerCase() === "delete") {
            DeleteDevice(deviceName).then(() => {
                setIsOpen(false);
            }).catch((error) => {
                console.error(error);
            });
        }
    }
    return <>
        <Modal isOpen={isOpen} onOpenChange={setIsOpen}>
            <ModalContent>
                <ModalHeader>
                    <span className="text-red-500">Are you sure you want to delete?</span>
                </ModalHeader>
                <ModalBody>
                    <p> <span className="font-extrabold">{deviceName}</span> will be deleted!</p>
                    <p>This action is irreversible.</p>
                    <p> Please type <span className="font-extrabold">DELETE</span> to confirm.</p>
                    <Input label="Confirmation" placeholder="Type DELETE to confirm" value={confirmation} onChange={(e) => setConfirmation(e.target.value)} />
                </ModalBody>
                <ModalFooter>
                    <Button color="danger" onClick={() => deleteDevice()} > Delete </Button>
                    <Button color="primary" onClick={() => setIsOpen(false)}> Cancel </Button>
                </ModalFooter>
            </ModalContent>
        </Modal >
    </>
}