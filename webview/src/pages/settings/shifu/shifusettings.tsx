import { AccordionItem, Accordion, Button, Checkbox, Select, SelectItem, Divider } from "@heroui/react";
import { useEffect, useState } from "react";
import { GetAllAvailableVersions, InstallShifu } from "src/apis/shifu/shifu";
import { ConfirmDelete } from "./confirmdelete";
import { useShifu } from "../../../hooks/useShifu";

export default function ShifuSettings() {
    const [versionList, setVersionList] = useState<string[]>(["latest"]);
    const [version, setVersion] = useState("latest");
    const [hiddenRCVersion, setHiddenRCVersion] = useState(true);
    const [loading, setLoading] = useState(false);
    const [isDeleteShifuOpen, setIsDeleteShifuOpen] = useState(false);

    const { isInstalled, checkInstallation } = useShifu();

    useEffect(() => {
        setLoading(true);
        GetAllAvailableVersions().then((versions) => {
            setVersionList(["latest", ...versions.filter((version) =>
                version.length > 0 && (!hiddenRCVersion || !version.includes("rc")))
            ]);
        }).finally(() => {
            setLoading(false);
        });
    }, [hiddenRCVersion]);

    useEffect(() => {
        checkInstallation();
    }, []);

    function installShifu() {
        setLoading(true);
        InstallShifu(version).then(() => {
            checkInstallation();
        }).catch((error) => {
            console.error(error);
        }).finally(() => {
            setLoading(false);
        });
    }

    return <div className="w-full p-2">
        <p className="text-3xl font-bold">Shifu</p>
        <div className="w-fit">
            <div className="flex flex-row items-center gap-4">
                <Select
                    defaultSelectedKeys={["latest"]}
                    variant="underlined"
                    label="Version"
                    className="w-32"
                    size="lg"
                >
                    {versionList.map((version) => (
                        <SelectItem
                            key={version}
                            textValue={version}
                            onPress={() => setVersion(version)}
                        >
                            {version}
                        </SelectItem>
                    ))}
                </Select>

                {isInstalled ?
                    <Button
                        color="danger"
                        onPress={() => setIsDeleteShifuOpen(true)}
                        isDisabled={loading}
                    >
                        Uninstall
                    </Button> :
                    <Button
                        color="primary"
                        onPress={() => installShifu()}
                        isDisabled={loading}
                    >
                        Install
                    </Button>
                }
            </div>
            <ConfirmDelete
                isOpen={isDeleteShifuOpen}
                setIsOpen={setIsDeleteShifuOpen}
                onUninstalled={checkInstallation}
            />
            <Divider orientation="vertical" />
            <div>
                <Accordion>
                    <AccordionItem key="1" title="Advanced Settings">
                        <Checkbox
                            isSelected={hiddenRCVersion}
                            onValueChange={setHiddenRCVersion}
                            isDisabled={loading}
                        >
                            Hide RC versions
                        </Checkbox>
                    </AccordionItem>
                </Accordion>
            </div>
        </div>
    </div>;
}   
