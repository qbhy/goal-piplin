import {FC, PropsWithChildren} from "react";
import classNames from "classnames";

type NavProps = {
    activeKey: string
} & PropsWithChildren
const Page: FC<NavProps> = ({activeKey, children}) => {
    return (
        <div className='flex'>
            <div className='text-base min-h-screen flex flex-col bg-gray-800'>
                <div className='flex-1'>
                    <div className='py-5 px-3 text-white font-bold'>Goal-Piplin</div>
                    <div
                        className={classNames('text-center p-3 hover:cursor-pointer', {
                            'text-white bg-gray-600': activeKey == 'console',
                            'text-gray-300': activeKey != 'console',
                        })}>
                        控制面板
                    </div>
                    <div
                        className={classNames('text-center p-3 hover:cursor-pointer', {
                            'text-white bg-gray-600': activeKey == 'activates',
                            'text-gray-300': activeKey != 'activates',
                        })}>
                        我的动态
                    </div>
                    <div
                        className={classNames('text-center p-3 hover:cursor-pointer', {
                            'text-white bg-gray-600': activeKey == 'notice',
                            'text-gray-300': activeKey != 'notice',
                        })}>
                        我的消息
                    </div>
                </div>
                <div className='text-center p-3 text-gray-300 hover:cursor-pointer'>退出</div>
            </div>
            <div className='flex-1'>
                {children}
            </div>
        </div>
    )
}

export default Page