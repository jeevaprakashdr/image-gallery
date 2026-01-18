import Gallery from "./components/Gallery";
import UploadImgForm from "./components/UploadImgForm";

export default function Home() {
  return (
    <div>
      <UploadImgForm />
      <Gallery/>
    </div>
  )
}
