import {FC, PropsWithChildren} from "react";
import classNames from "classnames";
import {getMyself} from "../services/auth.ts";
import {useRequest} from "ahooks";
import {Link} from "react-router-dom";

type NavProps = {
    activeKey: string
} & PropsWithChildren
const Page: FC<NavProps> = ({activeKey, children}) => {
    const {data} = useRequest(getMyself)
    return (
        <div className='flex'>
            <div className='text-base min-h-screen flex flex-col bg-gray-800'>
                <div className='flex-1 flex flex-col justify-between'>
                    <div>
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

                    <div className='mb-2'>
                        <Link to='/manage'
                            className={classNames('block text-center p-3 hover:cursor-pointer', {
                                'text-white bg-gray-600': activeKey == 'manage',
                                'text-gray-300': activeKey != 'manage',
                            })}>
                            管理后台
                        </Link>
                        <div
                            className={classNames('text-gray-300 text-center p-3 hover:cursor-pointer')}>
                            {data?.username || '加载中...'}
                        </div>
                    </div>
                </div>
                <div className='border-t border-t-gray-200 text-center p-3 text-gray-300 hover:cursor-pointer'>退出</div>
            </div>
            <div className='flex-1'>
                {children}
            </div>
        </div>
    )
}

export default Page