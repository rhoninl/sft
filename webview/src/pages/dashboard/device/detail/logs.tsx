import { Tab, Tabs } from "@heroui/react";
import { useState, useEffect, useRef, useCallback } from "react";
import { GetAllContainerName, GetDeviceShifuLogs } from "src/apis/shifu/device"
import { GetDeviceShifuLogsResponse } from "src/proto/proto/shifu/shifu_pb";

interface LogsProps {
    deviceName: string;
}

// Custom hook to manage container names and logs
function useDeviceLogs(deviceName: string) {
    const [containerNames, setContainerNames] = useState<string[]>([]);
    const [activeContainer, setActiveContainer] = useState<string>("");
    const [logs, setLogs] = useState<string[]>([]);

    // Fetch container names when deviceName changes
    useEffect(() => {
        async function fetchContainerNames() {
            try {
                const response = await GetAllContainerName(deviceName);
                setContainerNames(response.getContainerNamesList());
            } catch (error) {
                console.error("Error fetching container names:", error);
            }
        }
        fetchContainerNames();
    }, [deviceName]);

    // Auto-select the first container if none selected
    useEffect(() => {
        if (!activeContainer && containerNames.length > 0) {
            setActiveContainer(containerNames[0]);
        }
    }, [containerNames, activeContainer]);

    // Subscribe to logs when active container changes
    useEffect(() => {
        setLogs([]); // Reset logs on container change
        if (!activeContainer) return;

        const { stream, cancel } = GetDeviceShifuLogs(deviceName, activeContainer);

        stream.on("data", (response: GetDeviceShifuLogsResponse) => {
            setLogs(prevLogs => [...prevLogs, response.getLog()]);
        });
        stream.on("error", (error) => {
            console.error("Log stream error:", error);
        });
        stream.on("end", () => {
            console.info("Log stream ended.");
        });

        return () => {
            cancel();
        };
    }, [deviceName, activeContainer]);

    return { containerNames, activeContainer, setActiveContainer, logs };
}

export function Logs({ deviceName }: LogsProps) {
    const { containerNames, activeContainer, setActiveContainer, logs } = useDeviceLogs(deviceName);
    const logContainerRef = useRef<HTMLDivElement>(null);
    const [shouldAutoScroll, setShouldAutoScroll] = useState(true);

    // Auto-scroll to bottom
    const scrollToBottom = useCallback(() => {
        if (logContainerRef.current && shouldAutoScroll) {
            logContainerRef.current.scrollTop = logContainerRef.current.scrollHeight;
        }
    }, [shouldAutoScroll]);

    // Update auto-scroll flag on user scroll
    const handleScroll = () => {
        if (logContainerRef.current) {
            const { scrollTop, scrollHeight, clientHeight } = logContainerRef.current;
            setShouldAutoScroll(Math.abs(scrollHeight - clientHeight - scrollTop) < 10);
        }
    };

    // Scroll when logs update
    useEffect(() => {
        scrollToBottom();
    }, [logs, scrollToBottom]);

    return (
        <div>
            <Tabs
                onSelectionChange={(key) => setActiveContainer(key as string)}
                selectedKey={activeContainer}
            >
                {containerNames.map((name) => (
                    <Tab key={name} title={name}>
                        <div
                            ref={logContainerRef}
                            className="max-h-[500px] overflow-y-auto"
                            onScroll={handleScroll}
                        >
                            {logs.map((log, idx) => (
                                <div key={idx}>{log}</div>
                            ))}
                        </div>
                    </Tab>
                ))}
            </Tabs>
        </div>
    )
}