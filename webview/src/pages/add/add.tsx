import { Button, Divider, Drawer, DrawerBody, DrawerContent, DrawerFooter, DrawerHeader, Input } from "@heroui/react";
import CreateByUrl from "./url/url";
export default function Add({ isOpen, onClose }: { isOpen: boolean, onClose: () => void }) {
    console.log(isOpen);
    return (
        <Drawer isOpen={isOpen} onClose={onClose} size="xl" placement="left">
            <DrawerContent>
                <DrawerHeader>
                    Install Things in your Cluster
                </DrawerHeader>
                <DrawerBody>
                    <CreateByUrl />
                    <Divider />
                </DrawerBody>
                <DrawerFooter>
                    <Button color="primary" onPress={onClose}>Close</Button>
                </DrawerFooter>
            </DrawerContent>
        </Drawer>
    );
}