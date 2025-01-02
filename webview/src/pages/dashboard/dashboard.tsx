import './dashboard.css';
import ShifuInstall from '../../component/notice/shifui/installchecker';
import { Divider } from '@nextui-org/react';
import TopTab from '../../component/toptab/toptab';

export default function Dashboard() {
    return (
        <div className='w-full p-2'>
            <TopTab />
            <Divider className='mt-1' />
            <div className='flex flex-col gap-4 pt-5'>
                <ShifuInstall />
            </div>
        </div>
    )
}