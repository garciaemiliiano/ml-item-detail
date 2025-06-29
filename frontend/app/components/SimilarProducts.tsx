import Image from "next/image";
import { Star } from "lucide-react";
import { Card, CardContent } from "@/components/ui/card";

const similarProducts = [
  {
    id: 1,
    title: "iPhone 15 Pro 128GB Titanio Natural",
    price: 999999,
    originalPrice: 1199999,
    image: "/placeholder.svg?height=200&width=200",
    rating: 4.7,
    reviews: 1250,
    freeShipping: true,
  },
  {
    id: 2,
    title: "iPhone 14 Pro Max 256GB Morado Profundo",
    price: 899999,
    originalPrice: 1099999,
    image: "/placeholder.svg?height=200&width=200",
    rating: 4.6,
    reviews: 2100,
    freeShipping: true,
  },
  {
    id: 3,
    title: "Samsung Galaxy S24 Ultra 256GB Titanio",
    price: 1199999,
    originalPrice: 1399999,
    image: "/placeholder.svg?height=200&width=200",
    rating: 4.5,
    reviews: 890,
    freeShipping: true,
  },
  {
    id: 4,
    title: "iPhone 15 128GB Rosa",
    price: 799999,
    originalPrice: 899999,
    image: "/placeholder.svg?height=200&width=200",
    rating: 4.8,
    reviews: 1650,
    freeShipping: true,
  },
];

export default function SimilarProducts() {
  return (
    <div className="space-y-6">
      <h2 className="text-xl font-semibold">Productos similares</h2>

      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
        {similarProducts.map((product) => (
          <Card
            key={product.id}
            className="hover:shadow-lg transition-shadow cursor-pointer"
          >
            <CardContent className="p-4">
              <div className="space-y-3">
                <div className="aspect-square relative bg-gray-50 rounded-lg overflow-hidden">
                  <Image
                    src={product.image || "/placeholder.svg"}
                    alt={product.title}
                    fill
                    className="object-contain p-2"
                  />
                </div>

                <div className="space-y-2">
                  <h3 className="text-sm font-medium line-clamp-2 h-10">
                    {product.title}
                  </h3>

                  <div className="flex items-center gap-1">
                    <Star className="w-3 h-3 fill-yellow-400 text-yellow-400" />
                    <span className="text-xs">{product.rating}</span>
                    <span className="text-xs text-gray-500">
                      ({product.reviews})
                    </span>
                  </div>

                  <div className="space-y-1">
                    {product.originalPrice > product.price && (
                      <p className="text-xs text-gray-500 line-through">
                        ${product.originalPrice.toLocaleString()}
                      </p>
                    )}
                    <p className="text-lg font-bold">
                      ${product.price.toLocaleString()}
                    </p>
                  </div>

                  {product.freeShipping && (
                    <p className="text-xs text-green-600">Envío gratis</p>
                  )}
                </div>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
}
