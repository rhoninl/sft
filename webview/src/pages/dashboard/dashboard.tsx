import './dashboard.css';
import ShifuInstall from '../../components/notice/shifui/installchecker';
import { Divider } from "@heroui/react";
import TopTab from '../../components/toptab/toptab';
import { Routes, Route, useLocation, Navigate } from 'react-router-dom';
import Settings from '../settings/settings';
import DevicePage from './device/device';
import Device from './device/detail/detail';
import Terminal from '../terminal/terminal';
import Add from '../add/add';

export default function Dashboard() {
    const path = useLocation().pathname.split("/")[1];

    return (
        <div className='w-full dashboard dark:bg-gray-700 rounded-lg m-4'>
            <TopTab />
            <Divider />
            {path !== "settings" && (
                <div className='flex flex-col gap-4 pt-5'>
                    < ShifuInstall />
                </div>
            )}

            <Routes>
                <Route path="/devices" element={<DevicePage />} />
                <Route path="/telemetryservices" element={<div>Telemetry Services</div>} />
                <Route path="/settings" element={<Settings />} />
                <Route path="/terminal" element={<Terminal />} />
                <Route path="/devices/:name" element={<Device />} />
                <Route path="/" element={<Navigate to="/devices" />} />
            </Routes>
        </div >
    );
}
