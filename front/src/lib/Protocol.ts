export type AccountDetail = {
  accountId: string;
  name: string;
  introduction: string;
}

export type FreeSchedule = {
  scheduleId: string;
  from: string;
  to: string;
  unitPrice: number;
}

export type InvoiceDetail = {
  invoiceId: string;
  amount: number;
  paid: boolean;
}

export type PaymentState = 'unpaid' | 'processing' | 'complete' | 'error'

export type ErrorResponseJson = {
  message: string;
}
