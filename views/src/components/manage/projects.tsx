import {useState} from "react";
import classNames from "classnames";
import {useRequest} from "ahooks";
import {getProjects} from "../../services/projects.ts";

type Tab = 'project' | 'group'

export default function Projects() {
    const [tab, setTab] = useState<Tab>('project')
    const {data: projects} = useRequest(async (page: number = 1) => getProjects(page))

    return (
        <div className='px-3'>
            <div className='flex'>
                <div className={classNames('p-3 hover:cursor-pointer', {
                    'border-t-[3px] border-t-orange-600 border-x': tab == 'project',
                    'border-b': tab != 'project',
                })} onClick={() => setTab('project')}>项目管理
                </div>
                <div className={classNames('p-3 hover:cursor-pointer', {
                    'border-t-[3px] border-t-orange-600 border-x': tab == 'group',
                    'border-b': tab != 'group',
                })} onClick={() => setTab('group')}>项目分组
                </div>
                <div className='border-b flex-1'></div>
            </div>
            <div className='mt-3'>
                {
                    {
                        project: (
                            <div>
                                <div className='flex border-t-[3px] gap-x-5 py-3'>
                                    <div>ID</div>
                                    <div className='flex-1'>项目名称</div>
                                    <div className='flex-1'>代码仓库</div>
                                    <div>默认分支</div>
                                    <div>最近运行</div>
                                    <div>操作</div>
                                </div>
                                {projects?.list.map((project, index) => (
                                    <div className='flex gap-x-5 py-3' key={index}>
                                        <div>{project.id}</div>
                                        <div className='flex-1'>{project.name}</div>
                                        <div className='flex-1'>{project.repo_address}</div>
                                        <div>{project.default_branch}</div>
                                        <div>最近运行</div>
                                        <div>操作</div>
                                    </div>
                                ))}

                            </div>
                        ),
                        group: '',
                    }[tab]
                }
            </div>
        </div>
    )
}