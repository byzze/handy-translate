import React from "react";
import { RadioGroup, Radio } from "@nextui-org/react";
import { useEffect } from "react";
import toast, { Toaster } from 'react-hot-toast';
import { GetTransalteMap, GetTransalteWay, SetTransalteWay } from "../../../../../wailsjs/go/main/App"
import { useConfig, useToastStyle, useVoice, useSyncAtom } from '../../../../hooks';
import { atom, useAtom, useAtomValue } from 'jotai';

// import { translateServiceListAtom } from '../../../Translate';

export const translateServiceListAtom = atom([]);

export default function Way() {
    const [translateMap, setTranslateMap] = React.useState({});

    const [translateServiceList, setTranslateServiceList, syncTranslateServiceList] = useSyncAtom(translateServiceListAtom)

    // const [translateServiceList, setTranslateServiceList] = useAtomValue(translateServiceListAtom);


    const [selected, setSelected] = React.useState("");
    const toastStyle = useToastStyle();

    useEffect(() => {
        // GetTransalteWay().then(result => {
        //     setTranslateServiceList([result]);
        // });

        console.log(translateServiceList)
        GetTransalteMap().then(result => {
            result = JSON.parse(result)
            setTranslateMap(result)
        })

        GetTransalteWay().then(result => {
            setSelected(result)
        })
    }, [])


    return (
        <div className="flex flex-col gap-3">
            <Toaster />
            <RadioGroup
                label="选择你想要的翻译服务"
                value={selected}
                onValueChange={(value => {
                    console.log(value)
                    setSelected(value)
                    SetTransalteWay(value)
                    setTranslateServiceList([value])
                    timer = setTimeout(() => {
                        syncTranslateServiceList()
                    }, 1000);

                    console.log("test", translateServiceList)
                    toast.success('切换翻译' + value + '成功', {
                        style: toastStyle,
                    });
                })}
            >

                {translateMap &&
                    Object.entries(translateMap).map((key) => {
                        return (<Radio key={key[0]} value={key[0]}>{key[1].name}</Radio>)
                    })
                }
                {/* // <Radio value="sydney">Sydney</Radio>
                        // <Radio value="san-francisco">San Francisco</Radio>
                        // <Radio value="london">London</Radio>
                        // <Radio value="tokyo">Tokyo</Radio> */}
            </RadioGroup>
            <p className="text-default-500 text-small">Selected: {selected}</p>
        </div>
    );
}
