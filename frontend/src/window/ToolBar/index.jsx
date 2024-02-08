import React, { useEffect, useState, useRef } from "react";
import { Button, Card, CardBody, CardHeader, Divider } from "@nextui-org/react";
import { HeartIcon } from './HeartIcon';
import { CameraIcon } from './CameraIcon';
import { BsTranslate } from "react-icons/bs";
import { ToolBarShow, Show } from "../../../bindings/main/App";


export default function ToolBar() {
    const [result, setResult] = useState("")
    const textAreaRef = useRef();
    const [isLoading, setIsLoading] = useState(false)

    useEffect(() => {
        wails.Events.On("result", function (data) {
            let result = data.data

            setResult(result)
            let height = 0
            if (textAreaRef.current !== null) {
                textAreaRef.current.style.height = '0px';
                if (result !== '') {
                    // textAreaRef.current.scrollHeight 文本高度
                    height = textAreaRef.current.scrollHeight
                    textAreaRef.current.style.height = textAreaRef.current.scrollHeight + 'px';
                }
            }
            ToolBarShow(height)
            // ToolBarShow(height)
        })
    }, [])

    /*     useEffect(() => {
            if (textAreaRef.current !== null) {
                textAreaRef.current.style.height = '0px';
                if (result !== '') {
                    // textAreaRef.current.scrollHeight 文本高度
                    let height = textAreaRef.current.scrollHeight
                    textAreaRef.current.style.height = textAreaRef.current.scrollHeight + 'px';
                    ToolBarShow(height)
                }
            }
        }, [result]); */

    return (
        <div >
            <Card shadow='none'
                className='rounded-[10px]'>
                <CardHeader>
                    <div className="flex gap-4 items-center">
                        <Button size="sm" isIconOnly color="danger" aria-label="Like" onPress={() => {
                            Show("Translate")
                        }}>
                            <BsTranslate />
                        </Button>
                        {/* <Button size="sm" isIconOnly color="danger" aria-label="Like" onPress={() => {
                            Hide("Translate")
                        }}>
                            隐藏
                        </Button> */}
                    </div>
                </CardHeader>
                <Divider />
                <div className="flex w-full flex-wrap md:flex-nowrap mb-6 md:mb-0 gap-4">
                    <CardBody className={`p-[12px] pb-0 `}>
                        <textarea
                            ref={textAreaRef}
                            className='h-0 resize-none bg-transparent select-text outline-none'
                            readOnly
                            value={result}
                        />
                    </CardBody>
                </div>
            </Card >
        </div >
    );
}
