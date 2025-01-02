import { AccordionItem, Accordion, Button, Checkbox, Select, SelectItem } from "@nextui-org/react";
import { useEffect, useState } from "react";
import { GetAllAvailableVersions, InstallChecker, InstallShifu, UninstallShifu } from "src/apis/shifu/shifu";
import { ConfirmDelete } from "./confirmdelete";

export default function ShifuSettings() {
    const [versionList, setVersionList] = useState<string[]>(["latest"]);
    const [version, setVersion] = useState("latest");
    const [hiddenRCVersion, setHiddenRCVersion] = useState(true);
    const [isShifuInstalled, setIsShifuInstalled] = useState(false);
    const [loading, setLoading] = useState(false);
    const [isDeleteShifuOpen, setIsDeleteShifuOpen] = useState(false);
    const [statusChanged, setStatusChanged] = useState(false);

    useEffect(() => {
        setLoading(true);
        GetAllAvailableVersions().then((versions) => {
            setVersionList(["latest", ...versions.filter((version) => version.length > 0 && (!hiddenRCVersion || !version.includes("rc")))]);
        }).finally(() => {
            setLoading(false);
        });
    }, [hiddenRCVersion]);

    useEffect(() => {
        setLoading(true);
        InstallChecker().then((installed) => {
            setIsShifuInstalled(installed);
        }).finally(() => {
            setLoading(false);
        });
    }, []);

    useEffect(() => {
        if (statusChanged) {
            setLoading(true);
            InstallChecker().then((installed) => {
                setIsShifuInstalled(installed);
            }).finally(() => {
                setLoading(false);
                setStatusChanged(false);
            });
        }
    }, [statusChanged]);

    function installShifu() {
        setLoading(true);
        InstallShifu(version).then(() => {
            setStatusChanged(true);
        }).catch((error) => {
            console.error(error);
        }).finally(() => {
            setLoading(false);
        });
    }

    return <div className="w-full p-2">
        <p className="text-lg font-bold">Shifu</p>
        <div className="w-fit">
            <div className="flex flex-row items-center gap-4">
                <Select
                    defaultSelectedKeys={["latest"]}
                    variant="underlined"
                    label="Version"
                    className="w-32"
                >
                    {versionList.map((version) => (
                        <SelectItem
                            key={version}
                            value={version}
                            onClick={() => setVersion(version)}
                        >
                            {version}
                        </SelectItem>
                    ))}
                </Select>
                {isShifuInstalled ?
                    <Button color="danger" onClick={() => setIsDeleteShifuOpen(true)} isDisabled={loading}> Uninstall </Button> :
                    <Button color="primary" onClick={() => installShifu()} isDisabled={loading}> Install </Button>}
            </div>
            <ConfirmDelete isOpen={isDeleteShifuOpen} setIsOpen={setIsDeleteShifuOpen} setUninstalledShifu={setStatusChanged} />
            <div>
                <Accordion>
                    <AccordionItem key="1" title="Advanced Settings">
                        <Checkbox isSelected={hiddenRCVersion} onValueChange={setHiddenRCVersion} isDisabled={loading}> Hide RC versions</Checkbox>
                    </AccordionItem>
                </Accordion>
            </div>
        </div>
    </div>;
}   
