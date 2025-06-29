import {
  Star,
  Heart,
  Share2,
  Shield,
  Truck,
  CreditCard,
  ChevronRight,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import ImageGallery from "./ImageGallery";
import SellerInfo from "./SellerInfo";
import ProductSpecs from "./ProductSpecs";
import ReviewsSection from "./ReviewsSection";
import SimilarProducts from "./SimilarProducts";

async function getProductDetail(productId: string) {
  try {
    console.log("API URL:", process.env.NEXT_PUBLIC_API_URL);

    const response = await fetch(
      `http://backend:5000/items/${productId}`
    );

    console.log(response)
    if (!response.ok) {
      throw new Error("Failed to fetch product");
    }

    const data = (await response.json()).content;
    console.log("Product data:", data);
    return {
      id: data.id,
      title: data.product.name,
      price: data.product.price,
      image_url: data.product.image_url,
      description: data.product.description,
      stock: data.stock,
      inStock: data.stock > 0,
      fullWarranty: true,
      mercadoPago: true,
      seller: {
        name: data.provider.name,
      },
      category: {
        name: data.product.category.name,
      },
    };
  } catch (error) {
    console.error("Error fetching product:", error);
    return;
  }
}

export default async function ProductDetailPage() {
  const product = await getProductDetail(
    "ec24e26d-80ea-4069-bfcb-4205aef7dd2a"
  );
  return (
    <div className="min-h-screen bg-gray-50">
      <div className="bg-white border-b">
        <div className="max-w-7xl mx-auto px-4 py-3">
          <nav className="flex items-center space-x-2 text-sm text-gray-600">
            <span>Inicio</span>
            <ChevronRight className="w-4 h-4" />
            <ChevronRight className="w-4 h-4" />
            <span>{product?.category?.name}</span>
          </nav>
        </div>
      </div>

      <div className="max-w-7xl mx-auto px-4 py-6">
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <div className="lg:col-span-2">
            <ImageGallery image={product?.image_url} />
          </div>

          <div className="space-y-6">
            <div className="flex items-center gap-2"></div>

            <h1 className="text-2xl font-semibold text-gray-900 leading-tight">
              {product?.title}
            </h1>

            <div className="flex items-center gap-2">
              <div className="flex items-center">
                {[...Array(5)].map((_, i) => (
                  <Star
                    key={i}
                    // className={`w-4 h-4 ${
                    //   i < Math.floor(product.rating)
                    //     ? "fill-yellow-400 text-yellow-400"
                    //     : "text-gray-300"
                    // }`}
                  />
                ))}
              </div>
              {/* <span className="text-sm font-medium">{product.rating}</span>
              <span className="text-sm text-gray-600">
                ({product.reviewCount} opiniones)
              </span> */}
            </div>

            <div className="space-y-2">
              <div className="text-3xl font-bold text-gray-900">
                ${product?.price.toLocaleString()}
              </div>
            </div>

            <div className="flex items-center gap-2">
              <div className="w-2 h-2 bg-green-500 rounded-full"></div>
              <span className="text-sm text-green-600">Stock disponible</span>
              <span className="text-sm text-gray-600">
                ({product?.stock} disponibles)
              </span>
            </div>

            <div className="space-y-3">
              <div className="flex items-center gap-3">
                <Truck className="w-5 h-5 text-green-600" />
                <span className="text-sm">Envío gratis a todo el país</span>
              </div>
              <div className="flex items-center gap-3">
                <CreditCard className="w-5 h-5 text-blue-600" />
                <span className="text-sm">Pago seguro con Mercado Pago</span>
              </div>
            </div>

            <div className="space-y-3">
              <Button className="w-full bg-blue-600 hover:bg-blue-700 text-white py-3">
                Comprar ahora
              </Button>
              <Button variant="outline" className="w-full py-3">
                Agregar al carrito
              </Button>
            </div>

            <div className="flex items-center justify-between pt-4 border-t">
              <Button
                variant="ghost"
                size="sm"
                className="flex items-center gap-2"
              >
                <Heart className="w-4 h-4" />
                Agregar a favoritos
              </Button>
              <Button
                variant="ghost"
                size="sm"
                className="flex items-center gap-2"
              >
                <Share2 className="w-4 h-4" />
                Compartir
              </Button>
            </div>

            <SellerInfo seller={product?.seller.name} />
          </div>
        </div>

        <div className="mt-12">
          <Tabs defaultValue="description" className="w-full">
            <TabsList className="grid w-full grid-cols-4">
              <TabsTrigger value="description">Descripción</TabsTrigger>
              <TabsTrigger value="specs">Características</TabsTrigger>
              <TabsTrigger value="questions">Preguntas</TabsTrigger>
              <TabsTrigger value="reviews">Opiniones</TabsTrigger>
            </TabsList>

            <TabsContent value="description" className="mt-6">
              <Card>
                <CardContent className="p-6">
                  <div className="prose max-w-none">
                    <h3 className="text-lg font-semibold mb-4">
                      Descripción del producto
                    </h3>
                    <p className="text-gray-700 leading-relaxed mb-4">
                      {product?.description}
                    </p>
                    <p className="text-gray-700 leading-relaxed mb-4">
                      Características destacadas:
                    </p>
                    <ul className="list-disc list-inside space-y-2 text-gray-700"></ul>
                  </div>
                </CardContent>
              </Card>
            </TabsContent>

            <TabsContent value="specs" className="mt-6">
              <ProductSpecs />
            </TabsContent>

            <TabsContent value="questions" className="mt-6"></TabsContent>

            <TabsContent value="reviews" className="mt-6">
              {/* TODO <ReviewsSection
                rating={product.rating}
                reviewCount={product.reviewCount}
              /> */}
            </TabsContent>
          </Tabs>
        </div>
      </div>
    </div>
  );
}
