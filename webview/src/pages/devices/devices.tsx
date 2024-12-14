import './devices.css';
import { Divider } from '@nextui-org/react';
import ShifuInstall from '../../component/notice/shifui/installchecker';

export default function Devices() {
    return (
        <div className='w-full'>
            <h1 className='title mb-3'>Devices</h1>
            <ShifuInstall />
            <Divider className='mt-5' />
            <div className='flex flex-col gap-4 pt-5'>
                <a href="/devices/1">Device 1</a>
                <a href="/devices/2">Device 2</a>
                <a href="/devices/3">Device 3</a>
            </div>
        </div>
    )
}
