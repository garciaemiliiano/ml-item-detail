import { Star, MapPin } from "lucide-react"
import { Card, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"

interface SellerInfoProps {
  seller: {
    name: string
  }
}

export default function SellerInfo({ seller }: SellerInfoProps) {
  return (
    <Card>
      <CardContent className="p-4">
        <div className="space-y-3">
          <div className="flex items-center justify-between">
            <h3 className="font-semibold">Vendido por</h3>
            <Badge variant="secondary" className="bg-yellow-100 text-yellow-800">
   
            </Badge>
          </div>

          <div>
            <p className="font-medium text-blue-600">{seller?.name}</p>
            <div className="flex items-center gap-2 mt-1">
              <div className="flex items-center">
                <Star className="w-4 h-4 fill-yellow-400 text-yellow-400" />
              </div>
              {/* <span className="text-sm text-gray-600">({seller.sales.toLocaleString()} ventas)</span> */}
            </div>
          </div>

          <div className="flex items-center gap-2 text-sm text-gray-600">
            <MapPin className="w-4 h-4" />
          </div>

          <Button variant="outline" size="sm" className="w-full">
            Ver más productos del vendedor
          </Button>
        </div>
      </CardContent>
    </Card>
  )
}
