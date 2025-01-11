import { useEffect, useState, useMemo } from "react";
import { useParams } from "react-router-dom";
import { GetDeviceDetails } from "src/apis/shifu/device";
import { GetDeviceDetailsResponse } from "src/proto/proto/shifu/shifu_pb";
import { Button, Divider, Input } from "@nextui-org/react";
import "./detail.css"
import { ForwardDrawer } from "./forward";

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

    const gatewaySettings = useMemo(() => {
        try {
            if (device?.getEdgedevice()?.getGateway() === "null") return null;
            const gateway = JSON.parse(device?.getEdgedevice()?.getGateway() || "");

            const baseSettings = [
                { label: "Protocol", value: gateway.protocol },
                { label: "Address", value: gateway.address }
            ];

            const otherSettings = Object.entries(gateway).reduce((acc: Array<{ label: string, value: any }>, [key, value]) => {
                if (key === 'protocol' || key === 'address') return acc;

                if (value && typeof value === 'object' && !Array.isArray(value)) {
                    const flattenedSettings = Object.entries(value).map(([subKey, subValue]) => ({
                        label: subKey,
                        value: Array.isArray(subValue) ? subValue.join(', ') : subValue
                    }));
                    return [...acc, ...flattenedSettings];
                }
                return acc;
            }, []);

            return [...baseSettings, ...otherSettings];
        } catch (e) {
            return null;
        }
    }, [device]);

    const deviceSettings = useMemo(() => {
        try {
            if (device?.getEdgedevice()?.getSetting() === "null") return null;
            const settings = JSON.parse(device?.getEdgedevice()?.getSetting() || "");

            return Object.entries(settings).reduce((acc: Array<{ label: string, value: any }>, [key, value]) => {
                if (value && typeof value === 'object' && !Array.isArray(value)) {
                    const flattenedSettings = Object.entries(value).map(([subKey, subValue]) => ({
                        label: subKey,
                        value: Array.isArray(subValue) ? subValue.join(', ') : subValue
                    }));
                    return [...acc, ...flattenedSettings];
                }
                return acc;
            }, []);
        } catch (e) {
            return null;
        }
    }, [device]);

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
                <div className="flex flex-row gap-4">
                    <div className="grid grid-cols-2 gap-4 w-full">
                        {deviceDetails.map((detail, index) => (
                            <div key={index} className="flex items-center mx-2">
                                <div className="min-w-[4rem] whitespace-nowrap font-medium">{detail.label}:</div>
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
                {device?.getEdgedevice()?.getSetting() !== "null" && deviceSettings && (
                    <>
                        <Divider className="my-4" />
                        <div className="flex flex-col gap-4">
                            <h2 className="text-xl font-semibold mb-2">Device Settings</h2>
                            <div className="grid grid-cols-2 gap-4">
                                {deviceSettings.map((setting, index) => (
                                    <div key={index} className="flex items-center mx-2">
                                        <div className="min-w-[8rem] whitespace-nowrap font-medium">{setting.label}:</div>
                                        <Input
                                            value={setting.value || '-'}
                                            className="flex-1"
                                            variant="bordered"
                                            size="sm"
                                            readOnly
                                            isDisabled
                                        />
                                    </div>
                                ))}
                            </div>
                        </div>
                    </>
                )}
                {gatewaySettings && (
                    <>
                        <Divider className="my-4" />
                        <div className="flex flex-col gap-4">
                            <h2 className="text-xl font-semibold mb-2">Gateway Settings</h2>
                            <div className="grid grid-cols-2 gap-4">
                                {gatewaySettings.map((setting, index) => (
                                    <div key={index} className="flex items-center mx-2">
                                        <div className="min-w-[8rem] whitespace-nowrap font-medium">{setting.label}:</div>
                                        <Input
                                            value={setting.value || '-'}
                                            className="flex-1 "
                                            variant="bordered"
                                            size="sm"
                                            readOnly
                                            isDisabled
                                        />
                                    </div>
                                ))}
                            </div>
                        </div>
                    </>
                )}
            </div>
            <ForwardDrawer deviceName={name || ""} isOpen={isOpen} setIsOpen={setIsOpen} />
        </div >
    )
}

