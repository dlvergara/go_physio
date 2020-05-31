package errors

// General Error Codes
const UNAUTHORIZED = 3048
const INVALID_JSON_BODY = 3044
const INVALID_API_CREDENTIALS = 3045
const INVALID_MERCHANT_KEY = 3046
const UNEXPECTED_ERROR = 3087
const IDEMPOTENT_REQUEST_CONFLICT = 3089
const INVALID_QUERY_PARAMETER = 3097
const UNSUPPORTED_MEDIA_TYPE = 3098
const METHOD_NOT_ALLOWED = 3099
const PATH_NOT_FOUND = 3100

// Merchant Direct Error Codes
const MC_BOX_UOM_DIM_REQUIRED = 3001
const MC_BOX_UOM_DIM_INVALID = 3003
const MC_BOX_UOM_WEIGHT_REQUIRED = 3004
const MC_BOX_UOM_WEIGHT_INVALID = 3005
const MC_BOX_WEIGHT_REQUIRED = 3006
const MC_BOX_LEGTH_REQUIRED = 3007
const MC_BOX_WIDTH_REQUIRED = 3008
const MC_BOX_HEIGHT_REQUIRED = 3009
const MC_BOX_ITEMS_REQUIRED = 3010
const MC_BOX_ITEM_SKU_REQUIRED = 3011
const MC_BOX_ITEM_COO_REQUIRED = 3012
const MC_BOX_ITEM_COO_INVALID = 3013
const MC_SERVICE_LEVEL_DOWNGRADE_NOT_ALLOWED = 3014
const MC_TRANSIT_METHOD_CODE_REQUIRED = 3015
const MC_TRANSIT_METHOD_CODE_INVALID = 3016
const MC_TRANSIT_METHOD_CODE_NOT_SUPPORTED = 3017
const MC_UNRECOGNIZED_SKU = 3018
const MC_INVALID_PACK_ID = 3019
const MC_DOCUMENTS_NOT_READY = 3020
const MC_DOCUMENT_DOES_NOT_EXIST = 3021
const MC_CONSIGNOR_ADDRESS_REQUIRED = 3022
const MC_INSURANCE_FLAG_INVALID = 3023
const MC_INSURANCE_VALUE_INVALID = 3024
const MC_LABEL_FORMAT_NOT_SUPPORTED = 3025
const MC_ITEM_IS_CANCELLED = 3026
const RETAILER_ORDER_NUMBER_MISMATCH = 3027
const OPERATION_FAILED_ORDER_DOES_NOT_EXIST = 3028
const ORDER_HAS_BEEN_CANCELLED = 3029
const ORDER_DOCUMENTATION_NOT_READY = 3030
const ORDER_IDENTIFIER_REQUIRED = 3031
const MC_CARRIER_CREDENTIAL_REQUIRED = 3032
const MC_CARRIER_CREDENTIAL_INVALID = 3033
const MC_CARRIER_CREDENTIAL_EXPIRED = 3034
const MC_BOX_INFO_REQUIRED = 3035
const MC_BOX_INFO_MUST_BE_ARRAY = 3036
const MC_PICKUP_ADDRESS_REQUIRED = 3037
const MC_QTY_BOXES_GREATER_ZERO = 3038
const MC_QTY_BOXES_MISMATCH = 3039
const MC_PARTIAL_PACK_US_LIMIT = 3040
const MC_SERVICE_LEVEL_CATEGORY_MISSING = 3041
const MC_QUANTITY_ALLOWANCE_EXCEEDED = 3042
const MC_QUANTITY_EXCEEDED = 3043
const MC_ORDER_IS_NOT_MC = 3092

// Checkout API Error Codes
const SHIPPING_COULD_NOT_BE_CALCULATED = 2027
const ITEMS_NOT_AVAILABLE_FOR_EXPORT = 3002
const CART_NOT_FOUND_OR_EXPIRED = 3047
const MERCHANT_CONTROL_NOT_ENABLED = 3049
const INVALID_MERCHANT_CURRENCY = 3050
const INVALID_SHIPPING_METHOD_SELECTED = 3051
const INVALID_DOMESTIC_SHIPPING_MERCHANT_CONTROL = 3052
const MISSING_MERCHANT_SHIPPING_METHOD = 3053
const DUPLICATED_MERCHANT_SHIPPING_METHOD = 3054
const INVALID_MERCHANT_SHIPPING_METHOD_SELECTED = 3055
const INVALID_PRODUCT_ID = 3056
const INVALID_PRODUCT_NAME = 3057
const INVALID_PRODUCT_PRICE = 3058
const INVALID_PRODUCT_QUANTITY = 3059
const INVALID_PRODUCT_TRANSIT_TIME = 3060
const INVALID_ADDRESS_LINE_1 = 3061
const INVALID_ADDRESS_LINE_2 = 3062
const INVALID_COUNTRY_CODE = 3063
const SHIPPING_COUNTRY_NOT_ENABLED = 3064
const BILLING_COUNTRY_NOT_ENABLED = 3065
const INVALID_COUPON_DATA = 3066
const INVALID_PRODUCT_RETAIL_PRICE = 3067
const INVALID_SHIPPING_COUNTRY_CODE = 3068
const INVALID_SHIPPING_POSTAL_CODE = 3069
const SOME_ITEMS_CANT_BE_INCLUDED = 3070
const INTRA_COUNTRY_SHIPPING_NOT_AVAILABLE = 3071
const SHIPPING_NOT_AVAILABLE = 3072
const ITEM_CLASSIFIED_HAZARDOUS = 3073
const ITEM_IMPORT_RESTRICTION = 3074
const INVALID_CUSTOMER_IP_ADDRESS = 3075
const INVALID_SHIPPING_METHOD_CODE = 3076
const INVALID_SHIPPING_METHOD_NAME = 3077
const INVALID_DUTIABLE_TRANSIT_COST = 3078
const INVALID_CHARGED_TRANSIT_COST = 3079
const EXCHANGE_RATES_NOT_AVAILABLE = 3080
const ITEM_UPDATE_NOT_SUPPORTED = 3081
const SHIPPING_COUNTRY_UPDATE_NOT_SUPPORTED = 3082
const INVALID_SELLER_KEY = 3083
const INVALID_STATE_OR_PROVINCE = 3084
const RETAILER_ORDER_NUMBER_ALREADY_EXISTS = 3085
const EMPTY_CART = 3086
const MERCHANT_CURRENCY_REQUIRED = 3088
const SHIPPING_ADDRESS_REQUIRED = 3090
const CHECKOUT_DUTY_ID_CONFLICT = 3091
const ITEM_LICENSE_FLAG = 3093

// Checkout API - Payment
const TRANSACTION_CANT_BE_COMPLETED = 3101
const PAYMENT_PROCESSING_ERROR = 3102
const REQUIRED_3D_SECURITY = 3103
const INVALID_CARD_NUMBER = 3104
const INVALID_CCV_VALUE = 3105
const INVALID_CARD_TYPE = 3106
const INVALID_CHECKOUT_INFO = 3107
const INVALID_SECURITY_DETAILS = 3108
const EXPIRED_CARD = 3109
const EXPIRED_TRANSACTION_TIME = 3110
const REQUIRED_TRANSACTION_ID = 3111
const INVALID_TRANSACTION_ID = 3112
const TRANSACTION_ALREADY_CAPTURED = 3113
const ORDER_TOTALS_NOT_AVAILABLE = 3114

// Checkout API - Payment Resolution
const RESOLUTION_CONTACT_BANK = 3201
const RESOLUTION_CONTACT_BANK_IF_PERSISTS = 3202
const RESOLUTION_CONTACT_FCB_IF_PERSISTS = 3203

// Cancel Pack Notification
const MC_PACK_ALREADY_SHIPPED = 3301
const MC_PACK_ALREADY_CANCELLED = 3302

// Connect API
const ProductNotFound = 6003
const DutyIdNotEnabled = 6006
const DutyIdRequired = 6007
const DutyIdNotCorrect = 6008
const DutyIdExpired = 6009
const DutyIdShouldBeSameSku = 6010
const SellerKeyRequired = 6014
const Hs6Invalid = 6016
const DutyIdProductLimit = 6017
const OneProductAtLeastRequired = 6018
const DutyIdNotFound = 6019
const ProductIdRequiredNotEmpty = 6020
const DutyIdHsMismatch = 6025

// Shipment billing
const DATE_INVALID_FORMAT = 7000
const DATE_INVALID_PERIOD_END = 7001
const LIMIT_NEGATIVE = 7002
const LIMIT_OVER = 7003
const LIMIT_INVALID = 7004
const OFFSET_INVALID = 7005
const OFFSET_NEGATIVE = 7006
const NO_RECORDS = 7007
const PERIOD_START_MISSING = 7008
const PERIOD_END_MISSING = 7009

// Cancel order
const ORDER_CANCEL_NOT_AVAILABLE = 3400
const ORDER_CANCEL_SHIPPED = 3401

// Cancel item
const CANCEL_ITEM_SHIPPED = 3500
const CANCEL_ITEM_NOT_IN_ORDER = 3501
const CANCEL_ITEM_QUANTITY = 3502
const CANCEL_ITEM_CANCELLED = 3503
const CANCEL_ITEM_LC_ERROR = 3504
const CANCEL_ITEM_QUANTITY_ZERO = 3505

// Tracking Update
const TrackingUpdateNotSupported = 8201
const TrackingUpdateNotAvailableYet = 8202
const TrackingUpdateQuantityGreaterEq1 = 8203
const TrackingNumberIsRequired = 8204
const InvalidProductQuantity = 8205
const TrackingUpdateCannotBeExecuted = 8206
const InvalidTrackingCarrier = 8207

// Order Status
const ORDER_NOT_FOUND = 8301
const CANCEL_ORDER_EXPORT_SHIPPED = 8302

// Refund item
const REFUND_ITEM_NOT_SHIPPED = 4000
const REFUND_ITEM_NOT_IN_ORDER = 4001
const REFUND_ITEM_ORDER_CANCELLED = 4002
const REFUND_ITEM_FAILED = 4003
const REFUND_ITEM_ALREADY_REFUNDED = 4004

var errorCodes = map[int]map[string]interface{}{
	1004: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid product Country of Origin",
	},
	SHIPPING_COULD_NOT_BE_CALCULATED: map[string]interface{}{
		"status": 400,
		"msg":    "The shipping costs could not be calculated",
	},
	MC_BOX_UOM_DIM_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Box dimension unit of measurement is required",
	},
	ITEMS_NOT_AVAILABLE_FOR_EXPORT: map[string]interface{}{
		"status": 400,
		"msg":    "The following item(s) are not currently available for export. Please remove the item(s) from your cart to proceed with your checkout",
	},
	MC_BOX_UOM_DIM_INVALID: map[string]interface{}{
		"status": 400,
		"msg":    "Box dimension unit of measurement is invalid",
	},
	MC_BOX_UOM_WEIGHT_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Box weight unit of measurement is required",
	},
	MC_BOX_UOM_WEIGHT_INVALID: map[string]interface{}{
		"status": 400,
		"msg":    "Box weight unit of measurement is invalid",
	},
	MC_BOX_WEIGHT_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Box weight is required",
	},
	MC_BOX_LEGTH_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Box length is required",
	},
	MC_BOX_WIDTH_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Box width is required",
	},
	MC_BOX_HEIGHT_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Box height is required",
	},
	MC_BOX_ITEMS_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Box items are required",
	},
	MC_BOX_ITEM_SKU_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Item SKU is required",
	},
	MC_BOX_ITEM_COO_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Item country of origin is required",
	},
	MC_BOX_ITEM_COO_INVALID: map[string]interface{}{
		"status": 400,
		"msg":    "Item country of origin is invalid",
	},
	MC_SERVICE_LEVEL_DOWNGRADE_NOT_ALLOWED: map[string]interface{}{
		"status": 403,
		"msg":    "The airway bill service level is lower than the service level selected at the time of checkout. Please update the service level on your airway bill and resubmit, or contact us for assistance",
	},
	MC_TRANSIT_METHOD_CODE_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Transit Method Code is required",
	},
	MC_TRANSIT_METHOD_CODE_INVALID: map[string]interface{}{
		"status": 400,
		"msg":    "Transit Method Code is invalid",
	},
	MC_TRANSIT_METHOD_CODE_NOT_SUPPORTED: map[string]interface{}{
		"status": 400,
		"msg":    "Transit Method Code is not supported",
	},
	MC_UNRECOGNIZED_SKU: map[string]interface{}{
		"status": 400,
		"msg":    "An unrecognized SKU was provided",
	},
	MC_INVALID_PACK_ID: map[string]interface{}{
		"status": 404,
		"msg":    "Invalid or missing Pack Notification ID",
	},
	MC_DOCUMENTS_NOT_READY: map[string]interface{}{
		"status": 409,
		"msg":    "Documents are not ready yet for this Pack Notification",
	},
	MC_DOCUMENT_DOES_NOT_EXIST: map[string]interface{}{
		"status": 404,
		"msg":    "The document does not exist",
	},
	MC_CONSIGNOR_ADDRESS_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Consignor Address is required",
	},
	MC_INSURANCE_FLAG_INVALID: map[string]interface{}{
		"status": 400,
		"msg":    "Insurance must be boolean",
	},
	MC_INSURANCE_VALUE_INVALID: map[string]interface{}{
		"status": 400,
		"msg":    "Insurance value must be numeric and greater than 0",
	},
	MC_LABEL_FORMAT_NOT_SUPPORTED: map[string]interface{}{
		"status": 400,
		"msg":    "The requested label format is not supported",
	},
	MC_ITEM_IS_CANCELLED: map[string]interface{}{
		"status": 409,
		"msg":    "You may not issue a Pack Notification for an item which has been cancelled from this order",
	},
	RETAILER_ORDER_NUMBER_MISMATCH: map[string]interface{}{
		"status": 409,
		"msg":    "The requested retailer_order_number does not match with the value provided at checkout",
	},
	OPERATION_FAILED_ORDER_DOES_NOT_EXIST: map[string]interface{}{
		"status": 409,
		"msg":    "The operation cannot be performed because the referenced order number does not exist",
	},
	ORDER_HAS_BEEN_CANCELLED: map[string]interface{}{
		"status": 409,
		"msg":    "The order has been cancelled",
	},
	ORDER_DOCUMENTATION_NOT_READY: map[string]interface{}{
		"status": 409,
		"msg":    "The order documentation is not ready for shipment",
	},
	ORDER_IDENTIFIER_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Retailer_order_number is required if fxcb_order_number is empty",
	},
	MC_CARRIER_CREDENTIAL_REQUIRED: map[string]interface{}{
		"status": 403,
		"msg":    "The carrier credential token is required",
	},
	MC_CARRIER_CREDENTIAL_INVALID: map[string]interface{}{
		"status": 403,
		"msg":    "The carrier credential token is not valid",
	},
	MC_CARRIER_CREDENTIAL_EXPIRED: map[string]interface{}{
		"status": 403,
		"msg":    "The carrier credential is expired or disabled",
	},
	MC_BOX_INFO_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Boxes_info is required",
	},
	MC_BOX_INFO_MUST_BE_ARRAY: map[string]interface{}{
		"status": 400,
		"msg":    "Boxes_info must be an array",
	},
	MC_PICKUP_ADDRESS_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Pickup address is required",
	},
	MC_QTY_BOXES_GREATER_ZERO: map[string]interface{}{
		"status": 400,
		"msg":    "The box quantity must be an integer greater than 0",
	},
	MC_QTY_BOXES_MISMATCH: map[string]interface{}{
		"status": 400,
		"msg":    "The box quantity does not match the box information",
	},
	MC_PARTIAL_PACK_US_LIMIT: map[string]interface{}{
		"status": 403,
		"msg":    "We do not support partial Pack Notifications for US orders where total combined item values are greater than $2500.00 USD. Please send a complete Pack Notification for this order or contact us for assistance",
	},
	MC_SERVICE_LEVEL_CATEGORY_MISSING: map[string]interface{}{
		"status": 400,
		"msg":    "The selected shipping method doesn't have service level category",
	},
	MC_QUANTITY_ALLOWANCE_EXCEEDED: map[string]interface{}{
		"status": 400,
		"msg":    "The quantity of the notification exceeds the allowed quantity for the SKU",
	},
	MC_QUANTITY_EXCEEDED: map[string]interface{}{
		"status": 400,
		"msg":    "The quantity of the notification exceeds the quantity of the original order for the SKU",
	},
	INVALID_JSON_BODY: map[string]interface{}{
		"status": 400,
		"msg":    "The request body is not a valid JSON string",
	},
	INVALID_API_CREDENTIALS: map[string]interface{}{
		"status": 401,
		"msg":    "Invalid API Credentials",
	},
	INVALID_MERCHANT_KEY: map[string]interface{}{
		"status": 401,
		"msg":    "Invalid Merchant Key",
	},
	CART_NOT_FOUND_OR_EXPIRED: map[string]interface{}{
		"status": 404,
		"msg":    "Cart does not exist or has expired",
	},
	UNAUTHORIZED: map[string]interface{}{
		"status": 401,
		"msg":    "Unauthorized",
	},
	MERCHANT_CONTROL_NOT_ENABLED: map[string]interface{}{
		"status": 403,
		"msg":    "Merchant Control order is not available or enabled",
	},
	INVALID_MERCHANT_CURRENCY: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid Merchant Currency",
	},
	INVALID_SHIPPING_METHOD_SELECTED: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid Shipping Method selected",
	},
	INVALID_DOMESTIC_SHIPPING_MERCHANT_CONTROL: map[string]interface{}{
		"status": 403,
		"msg":    "The Domestic Shipping Charge cannot be set for Merchant Control orders",
	},
	MISSING_MERCHANT_SHIPPING_METHOD: map[string]interface{}{
		"status": 400,
		"msg":    "Missing Merchant Shipping Method",
	},
	DUPLICATED_MERCHANT_SHIPPING_METHOD: map[string]interface{}{
		"status": 400,
		"msg":    "Duplicated Merchant Shipping Method",
	},
	INVALID_MERCHANT_SHIPPING_METHOD_SELECTED: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid Merchant Shipping Method selected",
	},
	INVALID_PRODUCT_ID: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing Product Id",
	},
	INVALID_PRODUCT_NAME: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing Product Name",
	},
	INVALID_PRODUCT_PRICE: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing Product Price",
	},
	INVALID_PRODUCT_QUANTITY: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing Product Quantity",
	},
	INVALID_PRODUCT_TRANSIT_TIME: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid Product Transit Time",
	},
	INVALID_ADDRESS_LINE_1: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid Address Line 1",
	},
	INVALID_ADDRESS_LINE_2: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid Address Line 2",
	},
	INVALID_COUNTRY_CODE: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing Country Code",
	},
	SHIPPING_COUNTRY_NOT_ENABLED: map[string]interface{}{
		"status": 403,
		"msg":    "The selected Shipping Country is not enabled",
	},
	BILLING_COUNTRY_NOT_ENABLED: map[string]interface{}{
		"status": 403,
		"msg":    "The selected Billing Country is not enabled",
	},
	INVALID_COUPON_DATA: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid coupon data",
	},
	INVALID_PRODUCT_RETAIL_PRICE: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing Product Retail Price",
	},
	INVALID_SHIPPING_COUNTRY_CODE: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing Shipping Country Code'",
	},
	INVALID_SHIPPING_POSTAL_CODE: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing Destination Postal Code",
	},
	SOME_ITEMS_CANT_BE_INCLUDED: map[string]interface{}{
		"status": 400,
		"msg":    "Due to regulations placed by the order's destination country, the following item(s) can not be included in this order",
	},
	INTRA_COUNTRY_SHIPPING_NOT_AVAILABLE: map[string]interface{}{
		"status": 403,
		"msg":    "FedEx Cross Border is unable to accommodate intra-country shipments at this time",
	},
	SHIPPING_NOT_AVAILABLE: map[string]interface{}{
		"status": 403,
		"msg":    "Shipping is not currently available for the selected destination country",
	},
	ITEM_CLASSIFIED_HAZARDOUS: map[string]interface{}{
		"status": 403,
		"msg":    "This item is classified as Hazardous and cannot be shipped",
	},
	ITEM_IMPORT_RESTRICTION: map[string]interface{}{
		"status": 403,
		"msg":    "Due to import regulations of the destination country",
	},
	INVALID_CUSTOMER_IP_ADDRESS: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid Customer IP Addres",
	},
	INVALID_SHIPPING_METHOD_CODE: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing Shipping Method Code",
	},
	INVALID_SHIPPING_METHOD_NAME: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing Shipping Method Name",
	},
	INVALID_DUTIABLE_TRANSIT_COST: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing Dutiable Transit Cost",
	},
	INVALID_CHARGED_TRANSIT_COST: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing Charged Transit Cost",
	},
	EXCHANGE_RATES_NOT_AVAILABLE: map[string]interface{}{
		"status": 400,
		"msg":    "Exchange Rates are not available for the selected currencies",
	},
	ITEM_UPDATE_NOT_SUPPORTED: map[string]interface{}{
		"status": 403,
		"msg":    "This type of order does not support item updates",
	},
	SHIPPING_COUNTRY_UPDATE_NOT_SUPPORTED: map[string]interface{}{
		"status": 403,
		"msg":    "The shipping country or territory cannot be updated in this order",
	},
	INVALID_SELLER_KEY: map[string]interface{}{
		"status": 401,
		"msg":    "Invalid Seller Key",
	},
	INVALID_STATE_OR_PROVINCE: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid or missing state or province",
	},
	RETAILER_ORDER_NUMBER_ALREADY_EXISTS: map[string]interface{}{
		"status": 409,
		"msg":    "The Retailer Order Number has already been registered",
	},
	EMPTY_CART: map[string]interface{}{
		"status": 400,
		"msg":    "Your cart is empty",
	},
	UNEXPECTED_ERROR: map[string]interface{}{
		"status": 500,
		"msg":    "An unexpected exception occurred while processing the request",
	},
	MERCHANT_CURRENCY_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Merchant Currency is required",
	},
	IDEMPOTENT_REQUEST_CONFLICT: map[string]interface{}{
		"status": 409,
		"msg":    "Idempotent request conflicts with previous request",
	},
	UNSUPPORTED_MEDIA_TYPE: map[string]interface{}{
		"status": 415,
		"msg":    "Unsupported Media Type",
	},
	INVALID_QUERY_PARAMETER: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid query parameter",
	},
	METHOD_NOT_ALLOWED: map[string]interface{}{
		"status": 405,
		"msg":    "Method Not Allowed",
	},
	PATH_NOT_FOUND: map[string]interface{}{
		"status": 404,
		"msg":    "The requested URL path does not exist",
	},
	SHIPPING_ADDRESS_REQUIRED: map[string]interface{}{
		"status": 400,
		"msg":    "Shipping Address is required",
	},
	CHECKOUT_DUTY_ID_CONFLICT: map[string]interface{}{
		"status": 400,
		"msg":    "At least one duty ID is missing, expired or incorrect",
	},
	MC_ORDER_IS_NOT_MC: map[string]interface{}{
		"status": 403,
		"msg":    "This operation is available only for Merchant Control orders",
	},
	ITEM_LICENSE_FLAG: map[string]interface{}{
		"status": 403,
		"msg":    "The customs department of the destination country requires a certain license and or documentation in order for shipment to be made",
	},
	TRANSACTION_CANT_BE_COMPLETED: map[string]interface{}{
		"status": 500,
		"msg":    "Your transaction cannot be completed at this time",
	},
	PAYMENT_PROCESSING_ERROR: map[string]interface{}{
		"status": 500,
		"msg":    "An error occurred while processing your payment. Please try again",
	},
	REQUIRED_3D_SECURITY: map[string]interface{}{
		"status": 500,
		"msg":    "Your transaction cannot be completed at this time",
	},
	INVALID_CARD_NUMBER: map[string]interface{}{
		"status": 400,
		"msg":    "The card number entered is invalid or incorrect. Please check that all information is entered correctly and resubmit your order",
	},
	INVALID_CCV_VALUE: map[string]interface{}{
		"status": 400,
		"msg":    "The CCV code entered is invalid or incorrect. Please check that all information is entered correctly and resubmit your order",
	},
	INVALID_CARD_TYPE: map[string]interface{}{
		"status": 400,
		"msg":    "The type of card selected is invalid. Please correct the card type and resubmit your order",
	},
	INVALID_CHECKOUT_INFO: map[string]interface{}{
		"status": 400,
		"msg":    "The information entered at checkout was invalid. Please check that all information is entered correctly and resubmit your order",
	},
	INVALID_SECURITY_DETAILS: map[string]interface{}{
		"status": 500,
		"msg":    "An error occurred with the payment processor. Your order cannot be completed at this time",
	},
	EXPIRED_CARD: map[string]interface{}{
		"status": 400,
		"msg":    "The card entered at checkout has expired. Please check that all information is entered correctly and resubmit your order",
	},
	EXPIRED_TRANSACTION_TIME: map[string]interface{}{
		"status": 403,
		"msg":    "The time allowed to accept payment has expired",
	},
	REQUIRED_TRANSACTION_ID: map[string]interface{}{
		"status": 400,
		"msg":    "The transaction ID is required. Please try again",
	},
	INVALID_TRANSACTION_ID: map[string]interface{}{
		"status": 400,
		"msg":    "The transaction ID is invalid. Please try again",
	},
	TRANSACTION_ALREADY_CAPTURED: map[string]interface{}{
		"status": 403,
		"msg":    "This transaction has already been captured",
	},
	ORDER_TOTALS_NOT_AVAILABLE: map[string]interface{}{
		"status": 400,
		"msg":    "The Order Totals are not available. Please review your Order and try again",
	},
	RESOLUTION_CONTACT_BANK: map[string]interface{}{
		"status": 400,
		"msg":    "Please contact your bank or credit card company to assist",
	},
	RESOLUTION_CONTACT_BANK_IF_PERSISTS: map[string]interface{}{
		"status": 400,
		"msg":    "If this error persists, please contact your bank or credit card company to assist",
	},
	RESOLUTION_CONTACT_FCB_IF_PERSISTS: map[string]interface{}{
		"status": 400,
		"msg":    "If this error persists, please submit a ticket to crossborder_payment@ftn.fedex.com or call FedEx Cross Border Support at +1.203.683.4894",
	},
	MC_PACK_ALREADY_SHIPPED: map[string]interface{}{
		"status": 409,
		"msg":    "You are attempting to cancel a Pack Notification for a shipment which has already been picked up by the carrier. Please contact us for assistance",
	},
	MC_PACK_ALREADY_CANCELLED: map[string]interface{}{
		"status": 409,
		"msg":    "The Pack Notification has already been cancelled",
	},
	6003: map[string]interface{}{
		"status": 400,
		"msg":    "Product does not exist",
	},
	6004: map[string]interface{}{
		"status": 400,
		"msg":    "Product is not classified",
	},
	6005: map[string]interface{}{
		"status": 400,
		"msg":    "Could not calculate duty rate",
	},
	DutyIdNotEnabled: map[string]interface{}{
		"status": 403,
		"msg":    "Duty ID is not enabled for your account",
	},
	DutyIdRequired: map[string]interface{}{
		"status": 400,
		"msg":    "Duty ID is required for your account",
	},
	DutyIdNotCorrect: map[string]interface{}{
		"status": 400,
		"msg":    "Duty ID is not correct",
	},
	DutyIdExpired: map[string]interface{}{
		"status": 400,
		"msg":    "Duty ID has expired",
	},
	DutyIdShouldBeSameSku: map[string]interface{}{
		"status": 400,
		"msg":    "Duty ID should be the same, for all items that have the same sku",
	},
	SellerKeyRequired: map[string]interface{}{
		"status": 400,
		"msg":    "Seller Key is required for your account",
	},
	Hs6Invalid: map[string]interface{}{
		"status": 400,
		"msg":    "The HS6 harmonized code is invalid",
	},
	DutyIdProductLimit: map[string]interface{}{
		"status": 403,
		"msg":    "Exceeded the limit of Duty ID products configured for your account",
	},
	OneProductAtLeastRequired: map[string]interface{}{
		"status": 400,
		"msg":    "At least one product must be provided",
	},
	DutyIdNotFound: map[string]interface{}{
		"status": 404,
		"msg":    "The Duty Id does not exist",
	},
	ProductIdRequiredNotEmpty: map[string]interface{}{
		"status": 400,
		"msg":    "Product ID is required and cannot be empty",
	},
	DutyIdHsMismatch: map[string]interface{}{
		"status": 400,
		"msg":    "The Duty ID is not valid for the given SKU. HS code mismatch. To correct this issue, generate a new DutyId for that SKU.",
	},
	DATE_INVALID_FORMAT: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid format for date, must be ISO8601",
	},
	DATE_INVALID_PERIOD_END: map[string]interface{}{
		"status": 400,
		"msg":    "Period end is before period start",
	},
	LIMIT_NEGATIVE: map[string]interface{}{
		"status": 400,
		"msg":    "Limit must be greater than zero",
	},
	LIMIT_OVER: map[string]interface{}{
		"status": 400,
		"msg":    "Limit cannot be greater than 500",
	},
	LIMIT_INVALID: map[string]interface{}{
		"status": 400,
		"msg":    "Limit invalid",
	},
	OFFSET_INVALID: map[string]interface{}{
		"status": 400,
		"msg":    "Offset invalid",
	},
	OFFSET_NEGATIVE: map[string]interface{}{
		"status": 400,
		"msg":    "Offset must be greater or equal than zero",
	},
	NO_RECORDS: map[string]interface{}{
		"status": 400,
		"msg":    "No records found, try another period",
	},
	PERIOD_START_MISSING: map[string]interface{}{
		"status": 400,
		"msg":    "Period start is required",
	},
	PERIOD_END_MISSING: map[string]interface{}{
		"status": 400,
		"msg":    "Period end is required",
	},
	ORDER_CANCEL_NOT_AVAILABLE: map[string]interface{}{
		"status": 403,
		"msg":    "Order cancelation is not available for this order",
	},
	ORDER_CANCEL_SHIPPED: map[string]interface{}{
		"status": 403,
		"msg":    "Cannot cancel an order that has been shipped or partially shipped",
	},
	CANCEL_ITEM_SHIPPED: map[string]interface{}{
		"status": 409,
		"msg":    "An item cannot be cancelled after it has been shipped",
	},
	CANCEL_ITEM_NOT_IN_ORDER: map[string]interface{}{
		"status": 400,
		"msg":    "This item does not exist in the order",
	},
	CANCEL_ITEM_QUANTITY: map[string]interface{}{
		"status": 400,
		"msg":    "Cannot cancel a quantity larger than is available on the order",
	},
	CANCEL_ITEM_CANCELLED: map[string]interface{}{
		"status": 400,
		"msg":    "This item has already been canceled or refunded",
	},
	CANCEL_ITEM_LC_ERROR: map[string]interface{}{
		"status": 409,
		"msg":    "Cannot recalculate landedcost",
	},
	CANCEL_ITEM_QUANTITY_ZERO: map[string]interface{}{
		"status": 400,
		"msg":    "Quantity must be an integer greater than 0",
	},
	8201: map[string]interface{}{
		"status": 403,
		"msg":    "Tracking Update is not supported in this order",
	},
	8202: map[string]interface{}{
		"status": 403,
		"msg":    "Tracking Update is not available yet for this order",
	},
	8203: map[string]interface{}{
		"status": 400,
		"msg":    "Product quantity must be greater than or equal to 1",
	},
	8204: map[string]interface{}{
		"status": 400,
		"msg":    "Tracking Number is required and cannot be empty",
	},
	8205: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid product quantity",
	},
	8206: map[string]interface{}{
		"status": 400,
		"msg":    "The requested Tracking Update cannot be executed",
	},
	8207: map[string]interface{}{
		"status": 400,
		"msg":    "Invalid carrier code",
	},
	8301: map[string]interface{}{
		"status": 404,
		"msg":    "The requested order does not exist",
	},
	REFUND_ITEM_NOT_IN_ORDER: map[string]interface{}{
		"status": 400,
		"msg":    "Cannot refund an item that does not exist on the order",
	},
	REFUND_ITEM_NOT_SHIPPED: map[string]interface{}{
		"status": 400,
		"msg":    "An item cannot be refunded before it has shipped",
	},
	REFUND_ITEM_ORDER_CANCELLED: map[string]interface{}{
		"status": 400,
		"msg":    "Cannot refund an item for an order that has been cancelled",
	},
	REFUND_ITEM_FAILED: map[string]interface{}{
		"status": 409,
		"msg":    "Refund Failed, Contact Fedex Cross Border",
	},
	REFUND_ITEM_ALREADY_REFUNDED: map[string]interface{}{
		"status": 400,
		"msg":    "Cannot refund an item that has already been refunded",
	},
	CANCEL_ORDER_EXPORT_SHIPPED: map[string]interface{}{
		"status": 409,
		"msg":    "Cannot cancel an order that has been shipped or partially shipped to a FedEx Cross Border distribution center",
	},
}

func GetErrorMessage(code int) string {
	if errorCodes[code] != nil {
		return errorCodes[code]["msg"].(string)
	}
	return ""
}
