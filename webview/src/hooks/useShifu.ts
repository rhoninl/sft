import { useSelector, useDispatch } from 'react-redux';
import { RootState, AppDispatch } from '../store/store';
import { 
  checkShifuInstallation, 
  fetchDevices, 
  updateDeviceStatus 
} from '../store/shifuSlice';

export const useShifu = () => {
  const dispatch = useDispatch<AppDispatch>();
  const shifuState = useSelector((state: RootState) => state.shifu);

  const checkInstallation = () => {
    return dispatch(checkShifuInstallation());
  };

  const loadDevices = () => {
    return dispatch(fetchDevices());
  };

  const updateDevice = (name: string, status: string) => {
    dispatch(updateDeviceStatus({ name, status }));
  };

  return {
    ...shifuState,
    checkInstallation,
    loadDevices,
    updateDevice,
  };
}; 