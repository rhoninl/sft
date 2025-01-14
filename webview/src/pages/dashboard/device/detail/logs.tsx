import { Tab, Tabs } from "@nextui-org/react";
import { useState, useEffect, useRef, useCallback } from "react";
import { GetAllContainerName, GetDeviceShifuLogs } from "src/apis/shifu/device"
import { GetDeviceShifuLogsResponse } from "src/proto/proto/shifu/shifu_pb";

interface LogsProps {
    deviceName: string;
}

export function Logs({ deviceName }: LogsProps) {
    const [containerName, setContainerName] = useState<string>("");
    const [containerNames, setContainerNames] = useState<string[]>([]);
    const [logs, setLogs] = useState<string[]>([]);
    const logContainerRef = useRef<HTMLDivElement>(null);
    const [shouldAutoScroll, setShouldAutoScroll] = useState(true);

    // Listen for scroll events to determine if auto-scroll should be enabled
    const handleScroll = () => {
        if (!logContainerRef.current) return;

        const { scrollTop, scrollHeight, clientHeight } = logContainerRef.current;
        const isScrolledToBottom = Math.abs(scrollHeight - clientHeight - scrollTop) < 10;
        setShouldAutoScroll(isScrolledToBottom);
    };

    // Auto scroll to bottom
    const scrollToBottom = useCallback(() => {
        if (logContainerRef.current && shouldAutoScroll) {
            logContainerRef.current.scrollTop = logContainerRef.current.scrollHeight;
        }
    }, [shouldAutoScroll]);

    useEffect(() => {
        GetAllContainerName(deviceName).then((response) => {
            setContainerNames(response.getContainerNamesList());
        });
    }, [deviceName]);

    useEffect(() => {
        if (containerName === "") return;
        const { stream, cancel } = GetDeviceShifuLogs(deviceName, containerName);
        stream.on("data", (response: GetDeviceShifuLogsResponse) => {
            setLogs(prevLogs => [...prevLogs, response.getLog()]);
        }).on("error", (error) => {
            console.log("error: ", error);
        }).on("end", () => {
            console.log("end");
        });
        return () => {
            console.log("cancel");
            cancel();
        };
    }, [deviceName, containerName]);

    // Scroll to bottom when logs update
    useEffect(() => {
        scrollToBottom();
    }, [logs, scrollToBottom]);

    useEffect(() => {
        setLogs([]);
        setShouldAutoScroll(true);
    }, [containerName]);

    return (
        <div>
            <Tabs onSelectionChange={(key) => setContainerName(key as string)}>
                {containerNames.map((containerName) => (
                    <Tab key={containerName} title={containerName}>
                        <div
                            ref={logContainerRef}
                            className="max-h-[500px] overflow-y-auto"
                            onScroll={handleScroll}
                        >
                            {logs.map((log, index) => (
                                <div key={index}>{log}</div>
                            ))}
                        </div>
                    </Tab>
                ))}
            </Tabs>
        </div>
    )
}