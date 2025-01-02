import { Divider } from "@nextui-org/react";
import ShifuSettings from "./shifu/shifusettings";

export default function Settings() {
    return <div className="w-full p-2">
        <ShifuSettings />
        <Divider className="my-5" />
    </div>
}
