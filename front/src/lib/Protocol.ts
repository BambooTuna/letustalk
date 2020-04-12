
export type MentorDetail = {
  mentorId: string;
  name: string;
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
