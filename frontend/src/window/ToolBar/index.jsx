import React, { useEffect, useState, useRef } from "react";
import { Button, Card, CardBody, CardHeader } from "@nextui-org/react";
import { HeartIcon } from './HeartIcon';
import { CameraIcon } from './CameraIcon';
import { BsTranslate } from "react-icons/bs";

export default function ToolBar() {
    const [result, setResult] = useState("")
    const textAreaRef = useRef();
    const [isLoading, setIsLoading] = useState(false)

    useEffect(() => {
        // wails.Events.On("loading", function (data) {
        //     setIsLoading(data.data == 'true')
        // })
        wails.Events.On("result", function (data) {
            let result = data.data
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
                    <div className="flex gap-4 items-center">
                        <Button isIconOnly color="danger" aria-label="Like" onPress={() => {
                            window.go.main.App.Show("translate")
                        }}>
                            显示
                        </Button>
                        <Button isIconOnly color="danger" aria-label="Like" onPress={() => {
                            window.go.main.App.Hide("translate")
                        }}>
                            隐藏
                        </Button>
                    </div>
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
