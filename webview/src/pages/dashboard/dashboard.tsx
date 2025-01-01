import './dashboard.css';
import ShifuInstall from '../../component/notice/shifui/installchecker';
import { Divider } from '@nextui-org/react';
import TopTab from '../../component/toptab/toptab';

export default function Dashboard() {
    return (
        <div className='w-full p-5'>
            <h1 className='title mb-3 text-2xl font-bold text-blue-600'>Devices</h1>
            <ShifuInstall />
            <Divider className='mt-5' />
            <div className='flex flex-col gap-4 pt-5'>
                <TopTab />
            </div>
        </div>
    )
}