import React, { useEffect, useState, useRef } from "react";
import { Button, Card, CardBody, CardHeader } from "@nextui-org/react";
import { HeartIcon } from './HeartIcon';
import { CameraIcon } from './CameraIcon';
import { BsTranslate } from "react-icons/bs";
import { WindowSetSize, WindowGetSize } from "../../../wailsjs/runtime/runtime";
import { EventsOn } from "../../../wailsjs/runtime/runtime";

export default function TinyHome() {
    const [result, setResult] = useState("")
    const textAreaRef = useRef();
    useEffect(() => {
        WindowGetSize().then((w, h) => {
            console.log(w, h)
        })
        WindowSetSize(160, 300)
    }, [])
    useEffect(() => {
        EventsOn("result", (result) => {
            setResult(result)
        })
    }, [])

    useEffect(() => {
        if (textAreaRef.current !== null) {
            textAreaRef.current.style.height = '0px';
            if (result !== '') {
                textAreaRef.current.style.height = textAreaRef.current.scrollHeight + 'px';
            }
        }
    }, [result]);

    return (
        <div >
            <Card>
                <CardHeader>
                    <Button isIconOnly color="danger" aria-label="Like">
                        <BsTranslate />
                    </Button>
                </CardHeader>
                <CardBody>
                    <textarea
                        ref={textAreaRef}
                        className='h-0 resize-none bg-transparent select-text outline-none'
                        readOnly
                        value={result}
                    />
                </CardBody>
            </Card>
        </div>
    );
}
