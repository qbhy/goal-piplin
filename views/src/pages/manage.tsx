import Page from "../components/page.tsx";
import {useState} from "react";
import classNames from "classnames";
import {useRequest} from "ahooks";
import {getMyself} from "../services/auth.ts";
import Projects from "../components/manage/projects.tsx";

type Tab = 'projects' | 'materials' | 'users';

export default function Manage() {
    const [tab, setTab] = useState<Tab>('projects')
    const tabs = {
        projects: '项目管理',
        materials: '物料管理',
        users: '用户管理',
    };
    const {data} = useRequest(getMyself)

    return <Page activeKey='manage'>
        <div className='flex bg-[#f5f5f5]'>
            <div className='px-8 min-h-screen'>
                <div className='text-2xl py-5 text-center'>管理后台</div>
                <div className='flex flex-col gap-y-3'>
                    <div className={classNames('hover:text-green-500 hover:cursor-pointer', {
                        'text-green-500': tab == 'projects',
                    })} onClick={() => setTab('projects')}>项目管理
                    </div>
                    <div className={classNames('hover:text-green-500 hover:cursor-pointer', {
                        'text-green-500': tab == 'materials',
                    })} onClick={() => setTab('materials')}>资源管理
                    </div>
                    <div className={classNames('hover:text-green-500 hover:cursor-pointer', {
                        'text-green-500': tab == 'users',
                    })} onClick={() => setTab('users')}>用户管理
                    </div>
                </div>
            </div>
            <div className='flex-1 bg-white'>
                <div className='border-b-[3px] border-b-gray-300 pb-3 m-3 text-base text-gray-800'>
                    <span>控制面板</span> <span className='mx-1 text-gray-300 text-sm'>&gt;</span> {tabs[tab]}
                </div>
                <div>
                    {
                        {
                            projects: <Projects/>,
                            materials: '',
                            users: '',
                        }[tab]
                    }
                </div>
            </div>
        </div>
    </Page>
}

