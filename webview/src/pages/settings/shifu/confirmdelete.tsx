import { Button, ModalBody, Modal, ModalContent, ModalFooter, ModalHeader, Input } from "@nextui-org/react";
import { useState } from "react";
import { UninstallShifu } from "src/apis/shifu/shifu";

interface ConfirmDeleteProps {
    isOpen: boolean;
    setIsOpen: (isOpen: boolean) => void;
    setUninstalledShifu: (uninstalledShifu: boolean) => void;
}

export function ConfirmDelete({ isOpen, setIsOpen, setUninstalledShifu }: ConfirmDeleteProps) {
    const [confirmText, setConfirmText] = useState("");

    function uninstallShifu() {
        UninstallShifu().then(() => {
            setUninstalledShifu(true);
            setIsOpen(false);
        }).catch((error) => {
            console.error(error);
        });
    }

    return <>
        <Modal isOpen={isOpen} onOpenChange={setIsOpen}>
            <ModalContent>
                <ModalHeader>
                    <p>Are you sure you want to delete Shifu?</p>
                </ModalHeader>
                <ModalBody>
                    <p className="text-red-500">This action will remove Shifu from your system.</p>
                    <p className="text-red-500">This action is irreversible.</p>
                    <br />
                    <p>Type <span className="font-bold">DELETE</span> to confirm.</p>
                    <Input placeholder="DELETE" value={confirmText} onChange={(e) => setConfirmText(e.target.value)} />
                </ModalBody>
                <ModalFooter>
                    <Button color="danger" onClick={() => uninstallShifu()} isDisabled={confirmText.toLowerCase() !== "delete"}> Uninstall </Button>
                    <Button color="primary" onClick={() => setIsOpen(false)}> Cancel </Button>
                </ModalFooter>
            </ModalContent>
        </Modal>
    </>
}