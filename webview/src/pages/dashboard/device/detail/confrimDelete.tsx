import { Button, ModalBody, Modal, ModalContent, ModalFooter, ModalHeader, Input, toast, addToast } from "@heroui/react";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { DeleteDevice } from "src/apis/shifu/device";

interface ConfirmDeleteProps {
    deviceName: string;
    isOpen: boolean;
    setIsOpen: (isOpen: boolean) => void;
}

export function ConfirmDelete({ deviceName, isOpen, setIsOpen }: ConfirmDeleteProps) {
    const [confirmation, setConfirmation] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        setConfirmation("");
    }, [isOpen]);

    function deleteDevice() {
        if (confirmation.toLowerCase() === "delete") {
            DeleteDevice(deviceName).then(() => {
                addToast({
                    title: "Device " + deviceName + " deleted successfully",
                    timeout: 3000,
                    shouldShowTimeoutProgress: true,
                    color: "success",
                });


                setIsOpen(false);
                navigate("/devices");
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
                    <Button color="primary" onPress={() => setIsOpen(false)}> Cancel </Button>
                    <Button color="danger" onPress={() => deleteDevice()} > Delete </Button>
                </ModalFooter>
            </ModalContent>
        </Modal >
    </>
}