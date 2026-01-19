import { Image } from "@/models/Image";
type Props = {
    photo: Image
}

export default function ImgContainer({photo}: Props) {
    return (
        <div className="">
            <img
                className="rounded-xl"
                src={`http://localhost:9000/images/scaled-${photo.id}.png`}
                alt={photo.title}
                width={280}
                height={250}>
            </img>
        </div>
    )
}