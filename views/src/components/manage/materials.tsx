import {FC, useState} from "react";
import classNames from "classnames";
import {useRequest} from "ahooks";
import {getGroups, Group} from "../../services/manage/groups.ts";
import Modal from "../modal.tsx";
import {useForm} from "react-hook-form";
import {createKey, getKeys, Key} from "../../services/manage/keys.ts";

type Tab = 'key' | 'server_group'

type KeyEditorProps = {
    defaultValue?: Key,
    onSubmit: (project: Key) => void,
    onClose: () => void
}

const KeyEditor: FC<KeyEditorProps> = ({defaultValue, onSubmit, onClose}) => {
    const {
        register,
        handleSubmit,
        formState: {errors/**/},
    } = useForm<Key>()


    return (
        <form onSubmit={handleSubmit(onSubmit)} className='block bg-white rounded-lg shadow'>
            <div className='py-3 text-base flex justify-between px-3 items-center'>
                <span>{defaultValue ? `编辑密钥 (${defaultValue.id})` : '新建密钥'}</span>
                <button className='p-2 hover:text-black hover:cursor-pointer' onClick={onClose}>x</button>
            </div>

            <div className='border-y px-3'>
                <div className='flex items-center py-3 w-[25rem]'>
                    <span>密钥名称</span>
                    <input className={classNames('ml-2 flex-1 focus:outline-0 rounded p-1', {
                        'border border-red-500': errors.name,
                    })} placeholder='输入密钥名称' {...register("name", {
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

export default function Materials() {
    const [tab, setTab] = useState<Tab>('key')
    const [keyModal, setKeyModal] = useState<{ value?: Key, visible: boolean }>({visible: false})
    const [groupModal, setGroupModal] = useState<{ value?: Group, visible: boolean }>({visible: false})
    const {data: keys} = useRequest(getKeys)
    const {data: groups} = useRequest(getGroups)


    return (
        <div className='px-3'>
            <div className='border-b-[3px] border-b-gray-300 pb-3 my-3 text-base text-gray-800 flex justify-between'>
                <div>
                    <span>控制面板</span> <span className='mx-1 text-gray-300 text-sm'>&gt;</span> 项目管理
                </div>
                <div>
                    {tab == 'key' ? (
                        <div onClick={() => setKeyModal({visible: true})}
                             className='py-2 px-3 bg-blue-500 rounded text-white hover:bg-blue-600 hover:cursor-pointer'>
                            创建密钥
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
                    'border-t-[3px] border-t-orange-600 border-x': tab == 'key',
                    'border-b': tab != 'key',
                })} onClick={() => setTab('key')}>密钥管理
                </div>
                <div className={classNames('p-3 hover:cursor-pointer', {
                    'border-t-[3px] border-t-orange-600 border-x': tab == 'server_group',
                    'border-b': tab != 'server_group',
                })} onClick={() => setTab('server_group')}>机柜管理
                </div>
                <div className='border-b flex-1'></div>
            </div>

            <div className='mt-3'>
                {
                    {
                        key: (
                            <div>
                                <div className='flex border-t-[3px] gap-x-5 py-3'>
                                    <div>ID</div>
                                    <div className='flex-1'>密钥名称</div>
                                    <div>操作</div>
                                </div>
                                {keys?.map((key, index) => (
                                    <div className='flex gap-x-5 py-3' key={index}>
                                        <div>{key.id}</div>
                                        <div className='flex-1'>{key.name}</div>
                                        <div>操作</div>
                                    </div>
                                ))}
                            </div>
                        ),
                        server_group: (
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

            <Modal visible={keyModal.visible}>
                <KeyEditor onSubmit={key => {
                    createKey(key).then(() => {
                        setKeyModal({visible: false})
                    }).catch(e => {
                        setKeyModal({visible: false})
                        alert(e.message)
                    })
                }} onClose={() => setKeyModal({visible: false})}/>
            </Modal>
            <Modal visible={groupModal.visible}>
                <KeyEditor onSubmit={group => {
                    console.log(group)
                }} onClose={() => setGroupModal({visible: false})}/>
            </Modal>
        </div>
    )
}