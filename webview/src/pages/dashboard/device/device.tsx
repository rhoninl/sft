import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom";
import { ListDevices } from "src/apis/shifu/device"
import { useShifu } from "../../../hooks/useShifu";
import {
    Table,
    TableHeader,
    TableColumn,
    TableBody,
    TableCell,
    getKeyValue,
    TableRow,
} from "@heroui/react";
import { Device } from "src/proto/proto/shifu/shifu_pb";

export default function DevicePage() {
    const [devices, setDevices] = useState<Device[]>([]);
    const { isInstalled } = useShifu();
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate();
    useEffect(() => {
        if (!isInstalled) {
            return;
        }

        setLoading(true);

        const fetchDevices = () => {
            ListDevices().then((devices) => {
                setDevices(devices)
            }).catch(() => {
                console.log("Failed to fetch devices");
                return
            }).finally(() => {
                setLoading(false);
            });
        };

        fetchDevices(); // Initial fetch

        const intervalId = setInterval(fetchDevices, 5000); // Fetch every 5 seconds

        return () => clearInterval(intervalId); // Cleanup on unmount
    }, [isInstalled])

    if (!isInstalled) {
        return <div>Shifu is not installed</div>
    }

    if (!loading && devices.length === 0) {
        return <div>No devices found</div>
    }

    return (
        <Table aria-label="Devices list" isStriped>
            <TableHeader>
                <TableColumn key="name">Name</TableColumn>
                <TableColumn key="protocol">Protocol</TableColumn>
                <TableColumn key="address">Address</TableColumn>
                <TableColumn key="status">Status</TableColumn>
                <TableColumn key="age">Age</TableColumn>
            </TableHeader>
            <TableBody items={devices}>
                {(device) => (
                    <TableRow key={device.getName()} onClick={() => navigate(`/devices/${device.getName()}`)}>
                        {(columnKey) => {
                            if (columnKey === "address") {
                                const address = device.getAddress();
                                return <TableCell>{address.startsWith("[P]") ? "Positive Device" : address}</TableCell>;
                            }
                            return <TableCell>{getKeyValue(device.toObject(), columnKey)}</TableCell>;
                        }}
                    </TableRow>
                )}
            </TableBody>
        </Table>
    )
}