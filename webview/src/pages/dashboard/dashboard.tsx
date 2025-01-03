import './dashboard.css';
import ShifuInstall from '../../component/notice/shifui/installchecker';
import { Divider } from '@nextui-org/react';
import TopTab from '../../component/toptab/toptab';
import { Routes, Route, useLocation } from 'react-router-dom';
import Settings from '../settings/settings';

export default function Dashboard() {
    const path = useLocation().pathname.split("/")[1];
    
    return (
        <div className='w-full p-2 dashboard'>
            <TopTab />
            <Divider />
            {path !== "settings" && (
                <div className='flex flex-col gap-4 pt-5'>
                    <ShifuInstall />
                </div>
            )}

            <Routes>
                <Route path="/devices" element={<div>Devices</div>} />
                <Route path="/telemetryservices" element={<div>Telemetry Services</div>} />
                <Route path="/settings" element={<Settings />} />
            </Routes>
        </div>
    );
}
