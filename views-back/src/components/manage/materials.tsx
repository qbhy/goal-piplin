import {useState} from "react";
import classNames from "classnames";
import {useRequest} from "ahooks";
import {getGroups, Group} from "../../services/manage/groups.ts";
import Modal from "../modal.tsx";
import {createKey, getKeys, Key} from "../../services/manage/keys.ts";
import KeyEditor from "../editors/key.tsx";
import {copy} from 'clipboard'
import GroupEditor from "../editors/group.tsx";

type Tab = 'key' | 'server_group'

export default function Materials() {
    const [tab, setTab] = useState<Tab>('key')
    const [keyModal, setKeyModal] = useState<{ value?: Key, visible: boolean }>({visible: false})
    const [groupModal, setGroupModal] = useState<{ value?: Group, visible: boolean }>({visible: false})
    const {data: keys} = useRequest(getKeys)
    const {data: groups} = useRequest(getGroups)
    const [showKey, setShowKey] = useState('')

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
                            创建机柜
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
                                    <div className='flex gap-x-5 p-3 hover:bg-blue-200 rounded' key={index}>
                                        <div>{key.id}</div>
                                        <div className='flex-1'>{key.name}</div>
                                        <div>
                                            <button onClick={() => setShowKey(key.public_key)}>查看</button>
                                        </div>
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

            <Modal visible={showKey != ''}>
                <div className='w-[60%] bg-white p-3 rounded'>
                    <div className='border-b p-5'>{showKey}</div>
                    <div className='flex mt-3 justify-evenly gap-3'>
                        <div onClick={() => copy(showKey)}
                             className='py-2 text-center px-5 rounded text-white bg-blue-500 hover:cursor-pointer'>复制
                        </div>
                        <div onClick={() => setShowKey('')}
                             className='py-2 text-center px-5 rounded border border-blue-500 text-blue-500 hover:cursor-pointer'>关闭
                        </div>
                    </div>
                </div>
            </Modal>

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
                <GroupEditor onSubmit={group => {
                    console.log(group)
                }} onClose={() => setGroupModal({visible: false})}/>
            </Modal>
        </div>
    )
}