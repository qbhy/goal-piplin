import Page from "../components/page.tsx";

export default function Dashboard() {
    return <Page activeKey='console'>
        <div className='border-b-[3px] border-b-gray-300 pb-3 m-3 text-base text-gray-800'>
            <span>控制面板</span> <span className='mx-1 text-gray-300 text-sm'>&gt;</span> 我的项目
        </div>
        <div>

        </div>
    </Page>
}

