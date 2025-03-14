package my.mildotdev.pesan.pesan_android_stub

import android.net.Uri
import com.github.miltsm.pesan_grpc_stubs.PesanGrpcKt
import com.google.protobuf.ByteString
import io.grpc.CallCredentials
import io.grpc.Grpc
import io.grpc.InsecureChannelCredentials
import io.grpc.Metadata
import io.grpc.Metadata.ASCII_STRING_MARSHALLER
import io.grpc.TlsChannelCredentials
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.asExecutor
import my.mildotdev.pesan.android_stub.BuildConfig
import java.util.concurrent.Executor
import javax.inject.Qualifier

@Qualifier
@Retention(AnnotationRetention.BINARY)
annotation class PesanService

// NOTE: android client just use this stub to make calls,
// instead letting them to have control over url, ssl & etc.
object PesanStub {
    fun getStub(): PesanGrpcKt.PesanCoroutineStub {
        val uri =
            Uri.parse("${BuildConfig.BACKEND_SCHEME}://${BuildConfig.BACKEND_HOST}:${BuildConfig.BACKEND_PORT}")
        val channel = let {
//            val creds = TlsChannelCredentials
//                .newBuilder()
//                .requireFakeFeature()
//                .build()
            val builder = Grpc.newChannelBuilderForAddress(uri.host, uri.port, InsecureChannelCredentials.create())
            // TODO: mTLS config; server knows client, client knows server
            if (uri.scheme == "https")
                builder.useTransportSecurity()
            else
                builder.usePlaintext()

            builder.executor(Dispatchers.IO.asExecutor()).build()
        }

        return PesanGrpcKt.PesanCoroutineStub(channel)
    }
}

// NOTE: for protected calls
class PesanCredential(private val accessToken: ByteString) : CallCredentials() {
    override fun applyRequestMetadata(
        requestInfo: RequestInfo?,
        appExecutor: Executor?,
        applier: MetadataApplier?
    ) {
        val headers = Metadata().apply {
            put(Metadata.Key.of("Authorization", ASCII_STRING_MARSHALLER), "Bearer $accessToken")
        }
        applier?.apply(headers)
    }
}