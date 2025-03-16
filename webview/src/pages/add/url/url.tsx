import { addToast, Button, Input, toast } from "@heroui/react";
import { useEffect, useState } from "react";
import { InstallViaURL } from "../../../apis/shifu/add";

interface CreateByUrlProps {
    onSuccess: () => void;
}

export default function CreateByUrl({ onSuccess }: CreateByUrlProps) {
    const [inputUrl, setInputUrl] = useState("");
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");


    function createByUrl() {
        setLoading(true);
        console.log(inputUrl);

        // regex to extract URL from either direct URL input or kubectl apply -f command
        const urlRegex = /\b(?:https?:\/\/|www\.)\S+/;
        const match = inputUrl.match(urlRegex);

        if (!match) {
            setError("Invalid URL");
            return;
        }

        setError("");
        const url = match[0].replace(/['"]+$/, '');

        InstallViaURL(url).then(() => {
            console.log("Installation successful");
            addToast({
                title: "Installation successful",
                timeout: 3000,
                shouldShowTimeoutProgress: true,
                color: "success",
            });
        }).catch((err) => {
            console.error("Installation failed:", err);
            setError("Installation failed: " + err.message);
        });

        setTimeout(() => {
            setLoading(false);
            onSuccess();
        }, 1000);
    }

    return (
        <>
            <h1 className="font-bold"> Deploy with your link</h1>
            <div className="flex flex-row gap-2 h-14">
                <Input
                    isClearable
                    onClear={() => {
                        setInputUrl("");
                        setError("");
                    }}
                    placeholder="https://your-link.com/xxx.yaml"
                    value={inputUrl}
                    errorMessage={error}
                    isInvalid={error !== ""}
                    onChange={(e) => {
                        setInputUrl(e.target.value);
                        setError("");
                    }} />
                <Button
                    color="primary"
                    onPress={createByUrl}
                    isDisabled={inputUrl === ""}
                    disabled={inputUrl === ""}
                    isLoading={loading}
                >Add</Button>
            </div>
        </>
    );
}