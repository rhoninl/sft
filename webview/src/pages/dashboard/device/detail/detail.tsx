import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { GetDeviceDetails } from "src/apis/shifu/device";
import { GetDeviceDetailsResponse } from "src/proto/proto/shifu/shifu_pb";
import { Button, Divider, Drawer, DrawerBody, DrawerContent, DrawerFooter, DrawerHeader, Input } from "@nextui-org/react";
import { IoIosArrowRoundForward } from "react-icons/io";
import "./detail.css"

export default function Device() {
    const { name } = useParams()
    const [device, setDevice] = useState<GetDeviceDetailsResponse | null>(null)
    const [isOpen, setIsOpen] = useState(false)

    useEffect(() => {
        if (!name) return
        GetDeviceDetails(name).then((res) => {
            setDevice(res)
        })
    }, [name])

    const deviceDetails = [
        { label: "Address", value: device?.getEdgedevice()?.getAddress() },
        { label: "Protocol", value: device?.getEdgedevice()?.getProtocol() },
        { label: "Status", value: device?.getEdgedevice()?.getStatus() },
        { label: "SKU", value: device?.getEdgedevice()?.getSku() },
        { label: "Age", value: device?.getEdgedevice()?.getAge() }
    ]

    return (
        <div className="flex flex-col w-full p-6 rounded-lg shadow-lg">
            <div className="flex items-center mb-2">
                <h1 className="text-3xl font-bold text-blue-600 dark:text-blue-400">{name}</h1>
                <div className={`ml-4 px-3 py-1 rounded-full text-sm font-medium ${device?.getEdgedevice()?.getStatus() === 'Running'
                    ? 'bg-green-100 text-green-800'
                    : 'bg-red-100 text-red-800'
                    }`}>
                    {device?.getEdgedevice()?.getStatus()}
                </div>
            </div>

            <div className="rounded-xl">
                <div className="flex flex-row gap-4 min-h-[200px]">
                    <div className="grid grid-cols-2 w-full">
                        {deviceDetails.map((detail, index) => (
                            <div key={index} className="flex items-center mx-2">
                                <div className="w-20 font-medium">{detail.label}:</div>
                                <Input
                                    value={detail.value || '-'}
                                    className="flex-1"
                                    variant="bordered"
                                    size="sm"
                                    readOnly
                                    isDisabled
                                />
                            </div>
                        ))}
                    </div>
                    <Divider orientation="vertical" className="h-auto" />
                    <div className="w-32">
                        <div className="flex flex-col gap-3">
                            <Button color="primary" disableRipple onClick={() => setIsOpen(true)}>Forward</Button>
                            <Button color="warning" disableRipple>Restart</Button>
                            <Button color="danger" disableRipple>Delete</Button>
                        </div>
                    </div>
                </div>
            </div>
            <Drawer isOpen={isOpen} onOpenChange={setIsOpen}>
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
            </Drawer>
        </div>
    )
}

function Forward() {
    console.log("Forward")
}