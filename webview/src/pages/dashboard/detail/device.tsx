import { Divider } from "@nextui-org/react";
import { useNavigate, useParams } from "react-router-dom";

export default function Device() {
    const { id } = useParams()
    const navigate = useNavigate()

    return (
        <div className="flex flex-col w-full p-5">
            <div className="flex flex-row gap-3 mb-3">
                <button onClick={() => navigate(-1)} className="title text-2xl font-bold text-blue-600">☜</button>
                <h1 className='title text-2xl font-bold text-blue-600'>{id}</h1>
            </div>
            <Divider className='' />
            <h1> Status : Running </h1>
            <h1> Address: 192.168.1.1 </h1>
            <h1> Protocol: HTTP </h1>
        </div>
    )
}
