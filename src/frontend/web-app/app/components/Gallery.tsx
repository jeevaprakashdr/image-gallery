import GalleryClient from "./GalleryClient";

type Props = {
    topic?: string | undefined
}

export default async function Gallery({ topic } : Props) {
    return <GalleryClient topic={topic} />;
}