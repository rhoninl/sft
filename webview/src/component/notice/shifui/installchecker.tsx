import { Alert, Button } from "@nextui-org/react";
import { useEffect, useState } from "react"

import "./installchecker.css"
import { InstallChecker, InstallShifu } from "../../../apis/shifu/checker";
import Loading from "../../loading/loading";

export default function ShifuInstallChecker() {

    var [shifuInstalled, setShifuInstalled] = useState(true)
    var [installing, setInstalling] = useState(false)


    useEffect(() => {
        InstallChecker().then((installed) => {
            setShifuInstalled(installed)
        })
    }, [shifuInstalled])

    return <>
        {!shifuInstalled &&
            <Alert color="warning" title="Shifu is not installed in this cluster, please install Shifu first." endContent={
                <Button className={`shifu-install-button ${installing ? 'installing' : ''}`}
                    color="warning"
                    isLoading={installing}
                    variant="flat"
                    onClick={installShifu}>
                    Install
                </Button>
            } />

        }
        {installing && <div className="z-50 fixed top-0 left-0 w-full h-full bg-white bg-opacity-50 flex justify-center items-center">
            <Loading />
        </div>}
    </>

    function installShifu(e: React.MouseEvent<HTMLButtonElement>) {
        setInstalling(true)

        // InstallShifu("latest").then(() => {
        //     InstallChecker().then((installed) => {
        //         setShifuInstalled(installed)
        //     }).catch((error) => {
        //         console.error("Failed to install Shifu:", error)
        //     }).finally(() => {
        //         setInstalling(false)
        //     })
        // })

        setTimeout(() => {
            setInstalling(false)
        }, 3000)
    }
}