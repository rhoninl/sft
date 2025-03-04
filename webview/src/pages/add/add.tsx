import { Button, Divider, Drawer, DrawerBody, DrawerContent, DrawerFooter, DrawerHeader, Input } from "@heroui/react";

export default function Add({ isOpen, onClose }: { isOpen: boolean, onClose: () => void }) {
    console.log(isOpen);
    return (
        <Drawer isOpen={isOpen} onClose={onClose} size="xl">
            <DrawerContent>
                <DrawerHeader>
                    Add Anything
                </DrawerHeader>
                <DrawerBody>
                    <h1 className="font-bold"> Deploy with your link</h1>
                    <div className="flex flex-row gap-2">
                        <Input placeholder="https://your-link.com/xxx.yaml" />
                        <Button color="primary" onPress={() => { console.log("add anythings") }}>Add</Button>
                    </div>
                    <Divider />
                </DrawerBody>
                <DrawerFooter>
                    <Button variant="light" color="danger" onPress={onClose}>Cancel</Button>
                    <Button color="primary" onPress={onClose}>Add</Button>
                </DrawerFooter>
            </DrawerContent>
        </Drawer>
    );
}