import Gallery from "@/app/components/Gallery"

type Props = {
    params: {
        term: string
    }
}

export async function generateMetadata({ params }: Props) {
    const { term } = await params;
    return {
        title: `Results for ${term}`
    }
}

export default async function SearchResults({ params }: Props) {
    const { term } = await params;
    return <Gallery topic={term} />
}