import { useEffect } from 'react';
import { useShifu } from '../../hooks/useShifu';
import { Card, Spinner } from '@nextui-org/react';

export default function Devices() {
  const { 
    devices, 
    deviceLoading, 
    deviceError, 
    loadDevices 
  } = useShifu();

  useEffect(() => {
    loadDevices();
  }, []);

  if (deviceLoading) {
    return <Spinner />;
  }

  if (deviceError) {
    return <div>Error: {deviceError}</div>;
  }

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      {devices.map((device) => (
        <Card key={device.name} className="p-4">
          <h3 className="text-lg font-bold">{device.name}</h3>
          <p className="text-sm">Status: {device.status}</p>
          {device.lastSeen && (
            <p className="text-xs text-gray-500">
              Last seen: {new Date(device.lastSeen).toLocaleString()}
            </p>
          )}
        </Card>
      ))}
    </div>
  );
} 