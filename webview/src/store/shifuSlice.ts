import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { InstallChecker } from '../apis/shifu/shifu';

interface DeviceStatus {
  name: string;
  status: string;
  lastSeen?: string;
}

interface ShifuState {
  isInstalled: boolean;
  isLoading: boolean;
  error: string | null;
  devices: DeviceStatus[];
  deviceLoading: boolean;
  deviceError: string | null;
}

const initialState: ShifuState = {
  isInstalled: false,
  isLoading: false,
  error: null,
  devices: [],
  deviceLoading: false,
  deviceError: null,
};

export const checkShifuInstallation = createAsyncThunk(
  'shifu/checkInstallation',
  async () => {
    const isInstalled = await InstallChecker();
    return isInstalled;
  }
);

export const fetchDevices = createAsyncThunk(
  'shifu/fetchDevices',
  async () => {
    // Replace with your actual API call
    const response = await fetch('/api/devices');
    const data = await response.json();
    return data;
  }
);

const shifuSlice = createSlice({
  name: 'shifu',
  initialState,
  reducers: {
    updateDeviceStatus: (state, action) => {
      const { name, status } = action.payload;
      const device = state.devices.find(d => d.name === name);
      if (device) {
        device.status = status;
        device.lastSeen = new Date().toISOString();
      }
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(checkShifuInstallation.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(checkShifuInstallation.fulfilled, (state, action) => {
        state.isLoading = false;
        state.isInstalled = action.payload;
      })
      .addCase(checkShifuInstallation.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.error.message || 'Failed to check installation';
      })
      .addCase(fetchDevices.pending, (state) => {
        state.deviceLoading = true;
        state.deviceError = null;
      })
      .addCase(fetchDevices.fulfilled, (state, action) => {
        state.deviceLoading = false;
        state.devices = action.payload;
      })
      .addCase(fetchDevices.rejected, (state, action) => {
        state.deviceLoading = false;
        state.deviceError = action.error.message || 'Failed to fetch devices';
      });
  },
});

export const { updateDeviceStatus } = shifuSlice.actions;
export default shifuSlice.reducer; 