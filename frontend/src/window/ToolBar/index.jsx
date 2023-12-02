import React, { useEffect, useState, useRef } from "react";
import { Button, Card, CardBody, CardHeader, Divider } from "@nextui-org/react";
import { HeartIcon } from './HeartIcon';
import { CameraIcon } from './CameraIcon';
import { BsTranslate } from "react-icons/bs";

export default function ToolBar() {
    const [result, setResult] = useState("")
    const textAreaRef = useRef();
    const [isLoading, setIsLoading] = useState(false)

    useEffect(() => {
        wails.Events.On("result", function (data) {
            let result = data.data
            setResult(result)
        })
    }, [])

    useEffect(() => {
        if (textAreaRef.current !== null) {
            textAreaRef.current.style.height = '0px';
            let height = 0
            if (result !== '') {
                // textAreaRef.current.scrollHeight 文本高度
                // 40 + 55 窗口空白区域+翻译图标区域
                height = textAreaRef.current.scrollHeight + 40 + 55
                textAreaRef.current.style.height = textAreaRef.current.scrollHeight + 'px';
                window.go.main.App.ToolBarShow(height)
            }
        }
    }, [result]);

    return (
        <div >
            <Card shadow='none'
                className='rounded-[10px]'>
                <CardHeader>
                    <div className="flex gap-4 items-center">
                        <Button size="sm" isIconOnly color="danger" aria-label="Like" onPress={() => {
                            window.go.main.App.Show("translate")
                        }}>
                            <BsTranslate />
                        </Button>
                        {/* <Button size="sm" isIconOnly color="danger" aria-label="Like" onPress={() => {
                            window.go.main.App.Hide("translate")
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
