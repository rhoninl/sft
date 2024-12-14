import { Divider } from "@nextui-org/react";
import { useNavigate, useParams } from "react-router-dom";

export default function Device() {
    const { id } = useParams()
    const navigate = useNavigate()

    return (
        <div className="flex flex-col gap-3 w-full">
            <div className="flex flex-row gap-3">
                <button onClick={() => navigate(-1)} className="title">☜</button>
                <h1 className='title'>{id}</h1>
            </div>
            <Divider className='mt-3' />
            <h1> Status : Running </h1>
            <h1> Address: 192.168.1.1 </h1>
            <h1> Protocol: HTTP </h1>
        </div>
    )
}
