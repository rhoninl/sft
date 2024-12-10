import React from 'react';

import './devices.css';
import { Alert, Divider } from '@nextui-org/react';
export default function Devices() {
    return (<div className='devices'>
        <h1 style={{ fontSize: '40px' }}>Devices</h1>
        <Divider />
        <div className='flex flex-col gap-4'>
            <a href="/devices/1">Device 1</a>
            <a href="/devices/2">Device 2</a>
            <a href="/devices/3">Device 3</a>
        </div>
    </div>)
}