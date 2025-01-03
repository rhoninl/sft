import { Alert, Button } from "@nextui-org/react";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useDispatch, useSelector } from 'react-redux';
import "./installchecker.css";
import Loading from "../../loading/loading";
import { RootState } from "../../../store/store";
import { checkShifuInstallation } from "../../../store/shifuSlice";
import { AppDispatch } from "../../../store/store";

export default function ShifuInstallChecker() {
    const [installing, setInstalling] = useState(false);
    const navigate = useNavigate();
    const dispatch = useDispatch<AppDispatch>();
    const { isInstalled, isLoading, error } = useSelector((state: RootState) => state.shifu);

    useEffect(() => {
        dispatch(checkShifuInstallation());
    }, [dispatch]);

    if (isLoading) {
        return <Loading />;
    }

    if (error) {
        return <Alert color="danger" title="Error" description={error} />;
    }

    return (
        <>
            {!isInstalled && (
                <Alert
                    color="warning"
                    title="Shifu is not installed"
                    description="Shifu is not installed in this cluster, please install Shifu first."
                    endContent={
                        <Button
                            className={`shifu-install-button ${installing ? 'installing' : ''}`}
                            color="warning"
                            isLoading={installing}
                            variant="flat"
                            onClick={() => navigate("/settings")}
                        >
                            Install
                        </Button>
                    }
                />
            )}
            {installing && (
                <div className="z-50 fixed top-0 left-0 w-full h-full bg-white bg-opacity-50 flex justify-center items-center">
                    <Loading />
                </div>
            )}
        </>
    );
}