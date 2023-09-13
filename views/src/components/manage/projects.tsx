import {FC, useState} from "react";
import classNames from "classnames";
import {useRequest} from "ahooks";
import {createProject, getProjects, Project} from "../../services/projects.ts";
import {createGroup, getGroups, Group} from "../../services/manage/groups.ts";
import Modal from "../modal.tsx";
import {useForm} from "react-hook-form";

type Tab = 'project' | 'group'

const ProjectEditor: FC<{
    defaultValue?: Project,
    onSubmit: (project: Project) => void,
    onClose: () => void
}> = ({defaultValue, onSubmit, onClose}) => {
    const {data: groups} = useRequest(getGroups)
    const {
        register,
        handleSubmit,
        formState: {errors/**/},
    } = useForm<Project>()


    return (
        <form onSubmit={handleSubmit(onSubmit)} className='block bg-white rounded-lg shadow'>
            <div className='py-3 text-base flex justify-between px-3 items-center'>
                <span>{defaultValue ? `编辑项目 (${defaultValue.id})` : '新建项目'}</span>
                <button className='p-2 hover:text-black hover:cursor-pointer' onClick={onClose}>x</button>
            </div>

            <div className='border-y px-3'>
                <div className='flex items-center py-3 w-[25rem]'>
                    <span>项目名称</span>
                    <input className={classNames('ml-2 flex-1 focus:outline-0 rounded p-1', {
                        'border border-red-500': errors.name,
                    })} placeholder='输入项目名称' {...register("name", {
                        required: true,
                        minLength: 2,
                        maxLength: 20
                    })} />
                </div>

                <div className='flex items-center py-3 w-[25rem]'>
                    <span>分组</span>
                    <select className={classNames('ml-2 flex-1 focus:outline-0 rounded p-1', {
                        'border border-red-500': errors.name,
                    })} placeholder='输入项目名称' {...register("group_id")} >
                        {(groups && groups.length > 0) ? groups?.map(group => (
                            <option value={group.id}>{group.name}</option>
                        )) : <option value={0}>未选择</option>}
                    </select>
                </div>

                <div className='flex items-center py-3 w-[25rem]'>
                    <span>密钥</span>
                    <select className={classNames('ml-2 flex-1 focus:outline-0 rounded p-1', {
                        'border border-red-500': errors.name,
                    })} placeholder='输入项目名称' {...register("key_id")} >
                        {groups?.map(group => (
                            <option value={group.id}>{group.name}</option>
                        ))}
                    </select>
                </div>

                <div className='flex items-center py-3 w-[25rem]'>
                    <span>仓库地址</span>
                    <input type='url' className={classNames('ml-2 flex-1 focus:outline-0 rounded p-1', {
                        'border border-red-500': errors.name,
                    })} placeholder='请输入仓库地址' {...register("repo_address", {
                        required: true,
                    })} />
                </div>

                <div className='flex items-center py-3 w-[25rem]'>
                    <span>项目路径</span>
                    <input type='url' className={classNames('ml-2 flex-1 focus:outline-0 rounded p-1', {
                        'border border-red-500': errors.name,
                    })} placeholder='请输入项目路径' {...register("project_path", {
                        required: true,
                    })} />
                </div>

                <div className='flex items-center py-3 w-[25rem]'>
                    <span>默认分支</span>
                    <input type='url' className={classNames('ml-2 flex-1 focus:outline-0 rounded p-1', {
                        'border border-red-500': errors.name,
                    })} placeholder='请输入默认分支' {...register("default_branch", {
                        required: true,
                    })} />
                </div>

            </div>

            <div className='p-3 flex justify-evenly gap-3'>
                <button className='inline-block text-center py-2 bg-blue-500 text-white px-5 rounded'
                        type='submit'>提交
                </button>
                <button onClick={onClose}
                        className='inline-block text-center py-2 border border-blue-500 text-blue-500 px-5 rounded'>取消
                </button>
            </div>
        </form>
    )
}

const GroupEditor: FC<{
    defaultValue?: Group,
    onSubmit: (project: Group) => void,
    onClose: () => void
}> = ({defaultValue, onSubmit, onClose}) => {
    const {
        register,
        handleSubmit,
        formState: {errors/**/},
    } = useForm<Group>()


    return (
        <form onSubmit={handleSubmit(onSubmit)} className='block bg-white rounded-lg shadow'>
            <div className='py-3 text-base flex justify-between px-3 items-center'>
                <span>{defaultValue ? `编辑分组 (${defaultValue.id})` : '新建分组'}</span>
                <button className='p-2 hover:text-black hover:cursor-pointer' onClick={onClose}>x</button>
            </div>

            <div className='border-y px-3'>
                <div className='flex items-center py-3 w-[25rem]'>
                    <span>分组名称</span>
                    <input className={classNames('ml-2 flex-1 focus:outline-0 rounded p-1', {
                        'border border-red-500': errors.name,
                    })} placeholder='输入分组名称' {...register("name", {
                        required: true,
                        minLength: 2,
                        maxLength: 20
                    })} />
                </div>
            </div>

            <div className='p-3 flex justify-evenly gap-3'>
                <button className='inline-block text-center py-2 bg-blue-500 text-white px-5 rounded'
                        type='submit'>提交
                </button>
                <button onClick={onClose}
                        className='inline-block text-center py-2 border border-blue-500 text-blue-500 px-5 rounded'>取消
                </button>
            </div>
        </form>
    )
}

export default function Projects() {
    const [tab, setTab] = useState<Tab>('project')
    const [projectModal, setProjectModal] = useState<{ value?: Project, visible: boolean }>({visible: false})
    const [groupModal, setGroupModal] = useState<{ value?: Group, visible: boolean }>({visible: false})
    const {data: projects} = useRequest(async (page: number = 1) => getProjects(page))
    const {data: groups, refresh: refreshGroup} = useRequest(getGroups)


    return (
        <div className='px-3'>
            <div className='border-b-[3px] border-b-gray-300 pb-3 my-3 text-base text-gray-800 flex justify-between'>
                <div>
                    <span>控制面板</span> <span className='mx-1 text-gray-300 text-sm'>&gt;</span> 项目管理
                </div>
                <div>
                    {tab == 'project' ? (
                        <div onClick={() => setProjectModal({visible: true})}
                             className='py-2 px-3 bg-blue-500 rounded text-white hover:bg-blue-600 hover:cursor-pointer'>
                            创建项目
                        </div>
                    ) : (
                        <div onClick={() => setGroupModal({visible: true})}
                             className='py-2 px-3 bg-blue-500 rounded text-white hover:bg-blue-600 hover:cursor-pointer'>
                            创建分组
                        </div>
                    )}
                </div>
            </div>

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
                        group: (
                            <div>
                                <div className='flex border-t-[3px] gap-x-5 py-3'>
                                    <div>ID</div>
                                    <div className='flex-1'>分组名称</div>
                                    <div>操作</div>
                                </div>
                                {groups?.map((group, index) => (
                                    <div className='flex gap-x-5 py-3' key={index}>
                                        <div>{group.id}</div>
                                        <div className='flex-1'>{group.name}</div>
                                        <div>操作</div>
                                    </div>
                                ))}
                            </div>
                        ),
                    }[tab]
                }
            </div>

            <Modal visible={projectModal.visible}>
                <ProjectEditor onSubmit={project => {
                    createProject(project).then(() => {
                        setProjectModal({visible: false})
                    }).catch(e => {
                        setProjectModal({visible: false})
                        alert(e.message)
                    })
                }} onClose={() => setProjectModal({visible: false})}/>
            </Modal>
            <Modal visible={groupModal.visible}>
                <GroupEditor onSubmit={group => {
                    createGroup(group).then(() => {
                        setGroupModal({visible: false})
                        refreshGroup()
                    }).catch(e => {
                        setGroupModal({visible: false})
                        alert(e.message)
                    })
                }} onClose={() => setGroupModal({visible: false})}/>
            </Modal>
        </div>
    )
}