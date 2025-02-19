import { Divider } from "@heroui/react";
import ShifuSettings from "./shifu/shifusettings";

export default function Settings() {
    return <div className="w-full p-2">
        <ShifuSettings />
        <Divider />
    </div>
}
