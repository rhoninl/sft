import { useEffect, useState } from "react"
import { ListDevices } from "src/apis/shifu/device"
import {
    Table,
    TableHeader,
    TableColumn,
    TableBody,
    TableCell,
    getKeyValue,
    TableRow,
} from "@nextui-org/react";
import { Device } from "src/proto/proto/shifu/shifu_pb";

export default function DevicePage() {
    const [devices, setDevices] = useState<Device[]>([]);

    useEffect(() => {
        const fetchDevices = () => {
            ListDevices().then((devices) => {
                setDevices(devices)
            })
        };

        fetchDevices(); // Initial fetch

        const intervalId = setInterval(fetchDevices, 5000); // Fetch every 5 seconds

        return () => clearInterval(intervalId); // Cleanup on unmount
    }, [])

    if (devices.length === 0) {
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
                    <TableRow key={device.getName()}>
                        {(columnKey) => <TableCell>{getKeyValue(device.toObject(), columnKey)}</TableCell>}
                    </TableRow>
                )}
            </TableBody>
        </Table>
    )
}